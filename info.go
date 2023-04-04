package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/ztrue/tracerr"
)

func printUserInfo() error {
	url := "https://server.duinocoin.com/v2/users/" + opts.Username

	if opts.Verbose {
		if opts.Color {
			fmt.Fprintln(color.Output, m("Fetching user info from"), url)
		} else {
			fmt.Println("Fetching user info from", url)
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return tracerr.Wrap(err)
	}
	defer resp.Body.Close()

	if opts.Verbose {
		if opts.Color {
			fmt.Fprintln(color.Output, w("Status:"), resp.Status)
		} else {
			fmt.Println("Status:", resp.Status)
		}
		fmt.Println()
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return tracerr.Wrap(err)
	}

	if !gjson.Get(string(b), "success").Bool() {
		return tracerr.New(gjson.Get(string(b), "message").String())
	}

	if opts.Color {
		fmt.Fprintln(color.Output, w("Username:"), gjson.Get(string(b), "result.balance.username").String())
		fmt.Fprintln(color.Output, w("Created:"), gjson.Get(string(b), "result.balance.created").String())
		fmt.Fprintln(color.Output, w("Last Login:"), time.Unix(gjson.Get(string(b), "result.balance.last_login").Int(), 0).UTC())
		fmt.Fprintln(color.Output, w("Trust Score:"), gjson.Get(string(b), "result.balance.trust_score").Int())
		fmt.Fprintln(color.Output, w("Warnings:"), gjson.Get(string(b), "result.balance.warnings").Int())
		fmt.Println()
		fmt.Fprintln(color.Output, w("Balance:"), gjson.Get(string(b), "result.balance.balance").Float(), "DUCO")
		fmt.Println()
		fmt.Fprintln(color.Output, y("== Verification"))
		fmt.Fprintln(color.Output, w("Verified:"), gjson.Get(string(b), "result.balance.verified").String())
		fmt.Fprintln(color.Output, w("Verified By:"), gjson.Get(string(b), "result.balance.verified_by").String())
		fmt.Fprintln(color.Output, w("Verified Date:"), time.Unix(gjson.Get(string(b), "result.balance.verified_date").Int(), 0).UTC())
		fmt.Println()
		fmt.Fprintln(color.Output, y("== Staking"))
		fmt.Fprintln(color.Output, w("Stake Amount:"), gjson.Get(string(b), "result.balance.stake_amount").Float())
		fmt.Fprintln(color.Output, w("Stake Date:"), time.Unix(gjson.Get(string(b), "result.balance.stake_date").Int(), 0).UTC())
		fmt.Println()
		fmt.Fprintln(color.Output, y("== Mining"))
		fmt.Fprintln(color.Output, w("Miners:"), len(gjson.Get(string(b), "result.miners").Array()))
		fmt.Println()
		gjson.Get(string(b), "result.miners").ForEach(func(key, value gjson.Result) bool {
			fmt.Fprintln(color.Output, y("Miner #%d", key.Int()+1))
			fmt.Fprintln(color.Output, w("Identifier:"), value.Get("identifier").String())
			fmt.Fprintln(color.Output, w("Hashrate:"), value.Get("hashrate").Float(), "H/s")
			fmt.Fprintln(color.Output, g("Accepted:"), value.Get("accepted").Int())
			fmt.Fprintln(color.Output, r("Rejected:"), value.Get("rejected").Int())
			fmt.Fprintln(color.Output, w("Last Share:"), value.Get("sharetime").Float(), "s")
			fmt.Fprintln(color.Output, w("Difficulty:"), value.Get("diff").Float())
			fmt.Fprintln(color.Output, w("Algorithm:"), value.Get("algorithm").String())
			fmt.Fprintln(color.Output, w("Pool:"), value.Get("pool").String())
			fmt.Fprintln(color.Output, w("Software:"), value.Get("software").String())
			fmt.Fprintln(color.Output, w("Thread ID:"), value.Get("threadid").String())
			fmt.Println()

			return true
		})
		fmt.Println()
		fmt.Fprintln(color.Output, y("== Transactions"))
		gjson.Get(string(b), "result.transactions").ForEach(func(key, value gjson.Result) bool {
			fmt.Fprintln(color.Output, value.Get("sender").String(), mb("=>"), value.Get("recipient").String())
			fmt.Fprintln(color.Output, w("Amount:"), value.Get("amount").String())
			fmt.Fprintln(color.Output, w("Memo:"), value.Get("memo").String())
			fmt.Fprintln(color.Output, w("Datetime:"), value.Get("datetime").String())
			fmt.Fprintln(color.Output, w("Hash:"), value.Get("hash").String())
			fmt.Fprintln(color.Output, w("ID:"), value.Get("id").String())
			fmt.Println()

			return true
		})
	} else {
		fmt.Println("Username:", gjson.Get(string(b), "result.balance.username").String())
		fmt.Println("Created:", gjson.Get(string(b), "result.balance.created").String())
		fmt.Println("Last Login:", time.Unix(gjson.Get(string(b), "result.balance.last_login").Int(), 0).UTC())
		fmt.Println("Trust Score:", gjson.Get(string(b), "result.balance.trust_score").Int())
		fmt.Println("Warnings:", gjson.Get(string(b), "result.balance.warnings").Int())
		fmt.Println()
		fmt.Println("Balance:", gjson.Get(string(b), "result.balance.balance").Float(), "DUCO")
		fmt.Println()
		fmt.Println("== Verification")
		fmt.Println("Verified:", gjson.Get(string(b), "result.balance.verified").String())
		fmt.Println("Verified By:", gjson.Get(string(b), "result.balance.verified_by").String())
		fmt.Println("Verified Date:", time.Unix(gjson.Get(string(b), "result.balance.verified_date").Int(), 0).UTC())
		fmt.Println()
		fmt.Println("== Staking")
		fmt.Println("Stake Amount:", gjson.Get(string(b), "result.balance.stake_amount").Float())
		fmt.Println("Stake Date:", time.Unix(gjson.Get(string(b), "result.balance.stake_date").Int(), 0).UTC())
		fmt.Println()
		fmt.Println("== Mining")
		fmt.Println("Miners:", len(gjson.Get(string(b), "result.miners").Array()))
		fmt.Println()
		gjson.Get(string(b), "result.miners").ForEach(func(key, value gjson.Result) bool {
			fmt.Printf("Miner #%d\n", key.Int()+1)
			fmt.Println("Identifier:", value.Get("identifier").String())
			fmt.Println("Hashrate:", value.Get("hashrate").Float(), "H/s")
			fmt.Println("Accepted:", value.Get("accepted").Int())
			fmt.Println("Rejected:", value.Get("rejected").Int())
			fmt.Println("Last Share:", value.Get("sharetime").Float(), "s")
			fmt.Println("Difficulty:", value.Get("diff").Float())
			fmt.Println("Algorithm:", value.Get("algorithm").String())
			fmt.Println("Pool:", value.Get("pool").String())
			fmt.Println("Software:", value.Get("software").String())
			fmt.Println("Thread ID:", value.Get("threadid").String())
			fmt.Println()

			return true
		})
		fmt.Println()
		fmt.Println("== Transactions")
		gjson.Get(string(b), "result.transactions").ForEach(func(key, value gjson.Result) bool {
			fmt.Println(value.Get("sender").String(), "=>", value.Get("recipient").String())
			fmt.Println("Amount:", value.Get("amount").String())
			fmt.Println("Memo:", value.Get("memo").String())
			fmt.Println("Datetime:", value.Get("datetime").String())
			fmt.Println("Hash:", value.Get("hash").String())
			fmt.Println("ID:", value.Get("id").String())
			fmt.Println()

			return true
		})
	}

	return nil
}
