### Run 
1) Register a new OAuth application to your github account: https://github.com/settings/developers
You will get the clientID & clientSecret there.
2) Run the server `go run *.go --client-id=<CLIENT_ID> --client-secret="CLIENT_SECRET"`

### Flow 
3) Visit `http://localhost:8080/`, click `Login with github`. This will make a request to the oAuth gateway(https://github.com/login/oauth/authorize) of github with clientID & redirect_uri as flags.
4) Now give the authorization permission. If everything goes well, authorization server will redirect to the given redirect_uri.
5) This redirected uri is handled by our server. We get the authorizationCode from request URL, & request to https://github.com/login/oauth/access_token by setting clientID, clientSecret & authz code to get the token.

### Using token
6) In the response, We will get the token. We then redirects to `/hello` endpoint (with `access_token` set), which is also handled by us.
7) In `/hello` endpoint, we just use that access-token to get the user details from `https://api.github.com/user`. 