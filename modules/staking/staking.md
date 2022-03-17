# PLUGCHAIN SDK GO

## STAKING MODULE

- [Query](#query)
    - [QueryValidators](#Validators) --QueryValidators
    - [QueryValidator](#Validator) --QueryValidator
    - [QueryValidatorDelegations](#q-ValidatorDelegations) --QueryValidatorDelegations
    - [QueryValidatorUnbondingDelegations](#ValidatorUnbondingDelegations) --QueryValidatorUnbondingDelegations
    - [QueryDelegation](#Delegation) --QueryDelegation
    - [QueryUnbondingDelegation](#q-UnbondingDelegation) --QueryUnbondingDelegation
    - [QueryDelegatorDelegations](#DelegatorDelegations) --QueryDelegatorDelegations
    - [QueryDelegatorUnbondingDelegations](#DelegatorUnbondingDelegations) --QueryDelegatorUnbondingDelegations
    - [QueryRedelegations](#Redelegations) --QueryRedelegations
    - [QueryDelegatorValidators](#DelegatorValidators) --QueryDelegatorValidators
    - [QueryDelegatorValidator](#DelegatorValidator) --QueryDelegatorValidator
    - [QueryHistoricalInfo](#HistoricalInfo) --QueryHistoricalInfo
    - [QueryPool](#Pool) --QueryPool
    - [QueryParams](#Params) --QueryParams
- [TX](#tx)
    - [CreateValidator](#create) --CreateValidator
    - [EditValidator](#edit) --EditValidator
    - [Delegate](#delegate) --Delegate
    - [Undelegate](#unlegate) --Undelegate
    - [BeginRedelegate](#begin) --BeginRedelegate


# realization

## Query<a name="query"></a><br/>

#### QueryValidators<a name="Validators"></a><br/>
>Query all validators
```go
resp, err := client.Staking.QueryValidators("", 1, 10)
```

#### QueryValidator<a name="Validator"></a><br/>
>queries validator info for given validator address.
```go
resp, err := client.Staking.QueryValidator("gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e")
```

#### QueryValidatorDelegations<a name="ValidatorDelegations"></a><br/>
>queries delegate info for given validator
```go
resp, err := client.Staking.QueryValidatorDelegations("gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e", 1, 10)
```

#### QueryValidatorUnbondingDelegations<a name="ValidatorUnbondingDelegations"></a><br/>
>queries unbonding delegations of a validator.
```go
resp, err := client.Staking.QueryValidatorUnbondingDelegations("gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e", 1, 10)
```

#### QueryDelegation<a name="Delegation"></a><br/>
>queries delegate info for given validator delegator pair.
```go
resp, err := client.Staking.QueryDelegation("gx1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nvx529g", "gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e")
```

#### QueryUnbondingDelegation<a name="UnbondingDelegation"></a><br/>
>queries unbonding info for given validator delegator pair.
```go
resp, err := client.Staking.QueryUnbondingDelegation("gx1yhf7w0sq8yn6gqre2pulnqwyy30tjfc4v08f3x", "gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e")
```

#### QueryDelegatorDelegations<a name="DelegatorDelegations"></a><br/>
>queries all delegations of a given delegator address.
```go
resp, err := client.Staking.QueryDelegatorDelegations("gx1yhf7w0sq8yn6gqre2pulnqwyy30tjfc4v08f3x", 1, 10)
```

#### QueryDelegatorUnbondingDelegations<a name="DelegatorUnbondingDelegations"></a><br/>
>queries all unbonding delegations of a given delegator address.
```go
resp, err := client.Staking.QueryDelegatorUnbondingDelegations("gx1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nvx529g", 1, 10)
```

#### QueryRedelegations<a name="Redelegations"></a><br/>
>queries redelegations of given address.
```go
resp, err := client.Staking.QueryRedelegations(staking.QueryRedelegationsReq{
	DelegatorAddr:    "gx1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nvx529g",
	SrcValidatorAddr: "",
	DstValidatorAddr: "",
	Page:             0,
	Size:             0,
})
```

#### QueryDelegatorValidators<a name="DelegatorValidators"></a><br/>
>queries all validators info for given delegator
```go
resp, err := client.Staking.QueryDelegatorValidators("gx1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nvx529g", 1, 10)
```

#### QueryDelegatorValidator<a name="DelegatorValidator"></a><br/>
>queries validator info for given delegator validator pair.
```go
resp, err := client.Staking.QueryDelegatorValidator("gx1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nvx529g", "gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e")
```

#### QueryHistoricalInfo<a name="HistoricalInfo"></a><br/>
>Query the historical information of a given height
```go
//get latestBlockHeight at first
status, err := client.Status(context.Background())
height := status.SyncInfo.LatestBlockHeight	
height -= 10
resp, err := client.Staking.QueryHistoricalInfo(height)
```

#### QueryPool<a name="Pool"></a><br/>
>Query current equity pool
```go
resp, err := client.Staking.QueryPool()
```

#### QueryParams<a name="Params"></a><br/>
>Query the current equity parameter information
```go
resp, err := client.Staking.QueryParams()
```


## TX<a name="tx"></a><br/>

#### CreateValidator<a name="create"></a><br/>
>Send the transaction application to become the verifier and entrust a certain number of plugs to the verifier.
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000uplugcn") //Fee
rate := types.MustNewDecFromStr("0.1")
maxRate := types.MustNewDecFromStr("0.1")
maxChangeRate := types.MustNewDecFromStr("0.01")
minSelfDelegation := types.OneInt()
value, _ := types.ParseDecCoin("1uplugcn")
req1 := staking.CreateValidatorRequest{
	Moniker:           "haha",
	Rate:              rate,
	MaxRate:           maxRate,
	MaxChangeRate:     maxChangeRate,
	MinSelfDelegation: minSelfDelegation,
	Pubkey:            "gxpub1addwnpepq2u6ptadw90873wuq39307e334vxtznyxmayprd2n4rpyp2sg209q7u2jnm",
	Value:             value,
}
res, err := client.Staking.CreateValidator(req1, baseTx)
```

#### EditValidator<a name="edit"></a><br/>
>Modify the parameters of validation, including commission rate, verifier node name and other description information.
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000uplugcn") //Fee
commissionRate := types.MustNewDecFromStr("0.1")
minSelfDelegation := types.OneInt()
req1 := staking.EditValidatorRequest{
	Moniker:           "haha",
	Identity:          "identity",
	Website:           "website",
	SecurityContact:   "abbccdd",
	Details:           "fadsfas",
	CommissionRate:    commissionRate,
	MinSelfDelegation: minSelfDelegation,
}
res, err := client.Staking.EditValidator(req1, baseTx)
```

#### Delegate<a name="delegate"></a><br/>
>Pass to the verification person
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000uplugcn") //Fee
amount, _ := types.ParseDecCoin("10000uplugcn")
delegateReq := staking.DelegateRequest{
	ValidatorAddr: "gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e",
	Amount:        amount,
}
res, err := client.Staking.Delegate(delegateReq, baseTx)
```

#### Undelegate<a name="unlegate"></a><br/>
>Release the entrustment pass from the verifier
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000uplugcn") //Fee
amount, _ := types.ParseDecCoin("500uplugcn")
undelegateReq := staking.UndelegateRequest{
	ValidatorAddr: "gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e",
	Amount:        amount,
}
res, err := client.Staking.Undelegate(undelegateReq, baseTx)
```

#### BeginRedelegate<a name="begin"></a><br/>
>Transfer part or all of a delegate from one verifier to another
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000uplugcn") //Fee
amount, _ := types.ParseDecCoin("500uplugcn")
redelegateReq := staking.BeginRedelegateRequest{
	ValidatorSrcAddress: "gxvaloper1xyu87sqeuv3wfqmnfqcq3vglpgg6dp4nwzyu2e",
	ValidatorDstAddress: "",
	Amount:              amount,
}
res, err = client.Staking.BeginRedelegate(redelegateReq, baseTx)
```
