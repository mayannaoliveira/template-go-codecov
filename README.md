<div align="center">
 <!-- <a href="" > 
 <img src=""/> 
 </a>  -->
 <a href="https://go.dev/" > 
 <img src="https://img.shields.io/badge/GoLand-000000.svg?&logo=GoLand&logoColor=white"/> 
 </a>
  <a href="https://docs.codecov.com/docs/quick-start" > 
 <img src="https://img.shields.io/badge/Codecov-F01F7A.svg?e&logo=Codecov&logoColor=white"/> 
 </a>
<a href="https://codecov.io/gh/mayannaoliveira/template-go-codecov" > 
 <img src="https://codecov.io/gh/mayannaoliveira/template-go-codecov/graph/badge.svg?token=7IPNYBPMJQ"/> 
 </a>
 <a href="https://github.com/mayannaoliveira/template-go-codecov/actions/workflows/codecov.yaml" > 
 <img src="https://github.com/mayannaoliveira/template-go-codecov/actions/workflows/codecov.yaml/badge.svg"/> 
 </a>
 </div>
 </br>

## Template Go + Codecov

:umbrella: The best way to keep your code save and covering. 

> [!NOTE]
> Codecov is an amazing tool if you want to know more about, please, check the [quickstart](https://docs.codecov.com/docs/quick-start).

#### How to make it works
Please check the steps above to run the project:
- Clone the repository `git clone`. 
- Create `cal.go`.
- Create `cal_test.go`
- Open terminal to install the dependencies `go mod init codecov-demo` and `go get github.com/stretchr/testify/assert`
- Run Go test in terminal `go test ./...`
- Check the result `OK`.
- Create a new Codecov Token Key in `Settings > Acess > Toke Key`.
- Create a folder in project repository `.gitignore > workflows`.
- Ajust GitHub `Actions > Setting > Secrets > Insert the CODECOV_TOKENcode`.
- Insert the Token Key as secret.
- Check if the action is woking.

> [!CAUTION]
> The token key is a secret and you can revoke it and create another one. Please, read the documantation to know more about [token key](https://docs.codecov.com/docs/codecov-tokens).

#### Codecov graphics

You can also embed these graphs in your own application, either via API or URL. To see this page:

- Navigate to the repository of your choice from the repository overview.
- Click the "Badges and Graphs" section in the sidebar
- Follow the tutorial [here](https://docs.codecov.com/docs/graphs#sunburst).

<div align="center">
 <a href="https://docs.codecov.com/docs/graphs#sunburst" > 
 <img src="https://codecov.io/gh/mayannaoliveira/template-go-codecov/graphs/sunburst.svg?token=7IPNYBPMJQ"/> 
 </a> 
 <a href="https://docs.codecov.com/docs/graphs#sunburst" > 
 <img src="https://codecov.io/gh/mayannaoliveira/template-go-codecov/graphs/tree.svg?token=7IPNYBPMJQ"/> 
 </a>
  <a href="https://docs.codecov.com/docs/graphs#sunburst" > 
 <img src="https://codecov.io/gh/mayannaoliveira/template-go-codecov/graphs/icicle.svg?token=7IPNYBPMJQ"/> 
 </a>
  <!-- <a href="" > 
 <img src=""/> 
 </a> -->
 </div>
 </br>








