# PLUGCHAIN SDK GO

## GOV MODULE

- [Query](#query)
    - [QueryProposal](#proposal) --QueryProposal
    - [QueryProposals](#proposals) --QueryProposals
    - [QueryVote](#q-vote) --QueryVote
    - [QueryVotes](#votes) --QueryVotes
    - [QueryParams](#params) --QueryParams
    - [QueryDeposit](#q-deposit) --QueryDeposit
    - [QueryDeposits](#deposits) --QueryDeposits
    - [QueryTallyResult](#result) --QueryTallyResult
- [TX](#tx)
    - [SubmitProposal](#submit) --SubmitProposal
    - [Deposit](#deposit) --Deposit
    - [Vote](#vote) --Vote


# realization

## Query<a name="query"></a><br/>

#### QueryProposal<a name="proposal"></a><br/>
>Detailed information on a single proposal
```go
rep, err := client.Gov.QueryProposal(6)
```

#### QueryProposals<a name="proposals"></a><br/>
>Filter proposals by status. If not filled in, all proposals will be filtered
```go
pro, err := client.Gov.QueryProposals("") //PROPOSAL_STATUS_DEPOSIT_PERIOD
```

#### QueryVote<a name="q-vote"></a><br/>
>Query the details of a vote
```go
vote, err := client.Gov.QueryVote(5, "gx1yhf7w0sq8yn6gqre2pulnqwyy30tjfc4v08f3x")
```

#### QueryVotes<a name="votes"></a><br/>
>Query proposal vote
```go
vote, err := client.Gov.QueryVotes(5)
```

#### QueryParams<a name="params"></a><br/>
>Query the parameters of the governance process
```go
vote, err := client.Gov.QueryParams("voting")
```

#### QueryDeposit<a name="q-deposit"></a><br/>
>Mortgage information from a mortgager in the proposal by proposing ID query
```go
resp, err := client.Gov.QueryDeposit(2, "gx1yhf7w0sq8yn6gqre2pulnqwyy30tjfc4v08f3x")
```

#### QueryDeposits<a name="deposits"></a><br/>
>Query all mortgage information in the proposal
```go
resp, err := client.Gov.QueryDeposits(2)
```

#### QueryTallyResult<a name="result"></a><br/>
>Query the vote counting result of the proposal
```go
vote, err := client.Gov.QueryTallyResult(5)
```

## TX<a name="tx"></a><br/>

#### SubmitProposal<a name="submit"></a><br/>
>Submit proposals with initial delegation. The title, description, type and mortgage of the proposal can be provided directly
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
proposalId, res, err := client.Gov.SubmitProposal(gov.SubmitProposalRequest{
	Title:          "Community Pool Spend",
	Description:    "Pay me some Atoms!",
	Type:           "Text",
	InitialDeposit: types.NewDecCoins(types.NewDecCoin("plug", types.NewInt(1000))),
}, baseTx)
```

#### Deposit<a name="deposit"></a><br/>
>Pledge tokens for a proposal
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
amount, err := types.ParseDecCoins("2000plug")
res, err := client.Gov.Deposit(gov.DepositRequest{
	ProposalId: 3,
	Amount:     amount,
}, baseTx)
```

#### Vote<a name="vote"></a><br/>
>Proposal voting
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
voteReq := gov.VoteRequest{
	ProposalId: 5,
	Option:     "VOTE_OPTION_YES",
}
res, err := client.Gov.Vote(voteReq, baseTx)
```
