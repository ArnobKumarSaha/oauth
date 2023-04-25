## Build

### On [Dex](https://github.com/dexidp/dex) repo 
The StaticClients section of examples/config-dev.yaml of dex repo should look like this : 
```yaml
staticClients:
- id: kube-bind
  redirectURIs:
  - 'http://127.0.0.1:5555/callback'
  name: 'Kube Bind'
  secret: ZXhhbXBsZS1hcHAtc2VjcmV0
```

- Run the dex server 
`make build`
`./bin/dex serve examples/config-dev.yaml`

### On this repo
```yaml
cd dex-example
go install -v ./...
dex-example
```


## Run


### Hit http://127.0.0.1:5555/login/

```yaml
http://127.0.0.1:5556/dex/auth
?client_id=kube-bind
&redirect_uri=http%3A%2F%2F127.0.0.1%3A5555%2Fcallback
&response_type=code
&scope=openid+profile+email&state=I+wish+to+wash+my+irish+wristwatch
```

### Login with examples
```yaml
   http://127.0.0.1:5556/dex/approval
   ?req=yaf44bgf2kotokwq6fkqhmj33
   &hmac=pgDMFJnqVrnQnyDwDZWqCUJ66WGn79kMrnt9X200kkA
```

### Grant access
```yaml
   http://127.0.0.1:5555/callback
   ?code=tghthlytpgxtxmia2fifkbehg
   &state=I+wish+to+wash+my+irish+wristwatch
```

ID Token:
eyJhbGciOiJSUzI1NiIsImtpZCI6IjhlYTAwZDQyZjI4MWU2MTI3NmU5MDg2ZjYwOTEyY2E3ZTVmNjk4ZTEifQ.eyJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjU1NTYvZGV4Iiwic3ViIjoiQ2cwd0xUTTROUzB5T0RBNE9TMHdFZ1J0YjJOciIsImF1ZCI6Imt1YmUtYmluZCIsImV4cCI6MTY4MjUwMzM3MiwiaWF0IjoxNjgyNDE2OTcyLCJhdF9oYXNoIjoicDFnaV9haFNSSjhCOXBWMUluY1BZZyIsImNfaGFzaCI6InN4T0REYXdweGFhd0RTMVVBRWZPSVEiLCJlbWFpbCI6ImtpbGdvcmVAa2lsZ29yZS50cm91dCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYW1lIjoiS2lsZ29yZSBUcm91dCJ9.kPOQ4cF4IRhi9lhspUyKhC30kZiVu7tqwmDjN0ZcttTq8xuzSAKLB-ndTCTsih6X6s4128PAybdROvn13o_F2sn3rPpQnIk5OfD6DnoG9emr-hEzYzLsE0ESb8HNQD4NZ1wINcSxWlRV9TCChXiAAz0lhhizawYsxj_660KzhewRuFw5e-gWW2GJai5CcaXgzME27qz_EPs80J48_xw_w4O0kqEWMjllWk4g9_ciM2H3sjnL_QqgvWWVQ3tSr7tsDLv3hQyR_U3L4EPhA-px-WNiGYgT0pE0YAdQbq_MspJs-_gf58JlkwjuwkkdYtgFK_VSUsNh6xDJQAQUcGs6kA

Access Token:
eyJhbGciOiJSUzI1NiIsImtpZCI6IjhlYTAwZDQyZjI4MWU2MTI3NmU5MDg2ZjYwOTEyY2E3ZTVmNjk4ZTEifQ.eyJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjU1NTYvZGV4Iiwic3ViIjoiQ2cwd0xUTTROUzB5T0RBNE9TMHdFZ1J0YjJOciIsImF1ZCI6Imt1YmUtYmluZCIsImV4cCI6MTY4MjUwMzM3MiwiaWF0IjoxNjgyNDE2OTcyLCJhdF9oYXNoIjoidndiOU1kTXlkQWx3ZTBmVVRVeVhXdyIsImVtYWlsIjoia2lsZ29yZUBraWxnb3JlLnRyb3V0IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiJLaWxnb3JlIFRyb3V0In0.pmsIiq6Fx2CMHizInLouyKGu81jKSJsBklKoBrFXaKbKqfHXnHMWoZFH4RQAk7FskAfFL1I5YtGl_wVyzObXgGp1WSzGUWmfYZWnG2R6ngXW6IsG7gWve2OTcjbzsAWaiJebBVVZCd05T-8I3sEZcT50yH_EPxAkngWYu6P_XPE2G_4w5QnVI7yZ-1iQNKCPzpOidvUwvIe6-29A067nh-GsSgAIqWE1qqWqVNCgaj42u57IZdbVLThK7Zn0wzRe_sg6xrxpTHn68iZiaVXx51cHGz_uAJiH5xkuZ2nMZSOhHprAzG-60TXWRXR4A7k33fJ6hZB_Zz4fRIZfx0a4Pw

Claims:
{
"iss": "http://127.0.0.1:5556/dex",
"sub": "Cg0wLTM4NS0yODA4OS0wEgRtb2Nr",
"aud": "kube-bind",
"exp": 1682503372,
"iat": 1682416972,
"at_hash": "p1gi_ahSRJ8B9pV1IncPYg",
"c_hash": "sxODDawpxaawDS1UAEfOIQ",
"email": "kilgore@kilgore.trout",
"email_verified": true,
"name": "Kilgore Trout"
}