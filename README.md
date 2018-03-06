# Databricks SDK for Go

databricks-sdk-go is an unofficial Databricks SDK for the Go programming language. It currently only supports a subset of the operations provided by the [Databricks REST API](https://docs.databricks.com/api/latest/index.html). At this point, there is only support for the following APIs:

* Workspace API

## Installation

To install the SDK run the following command:

```bash
go get -u github.com/betabandido/databricks-sdk-go
```

## Usage

### Creating an API Client

All operations interacting with Databricks REST API need an API Client. The following example shows how to create a client.

```golang
cl, err := client.NewClient(domain, token)
if err != nil {
    panic(err)
}
```

`domain` refers to the name of the Databricks deployment (e.g., `<your-account>.cloud.databricks.com`), and `token` needs to contain the authentication token. See Databricks [authentication documentation](https://docs.databricks.com/api/latest/examples.html) for more details.

### Making Requests

Requests for a given API can be sent using the appropriate endpoint. For instance, the following example shows how to upload a notebook to the workspace.

```golang
language := models.PYTHON
content := base64.StdEncoding.EncodeToString([]byte("print('hello world')"))

err := endpoint.Import(&models.WorkspaceImportRequest{
    Path:     path,
    Language: &language,
    Content:  content,
})
if err != nil {
    panic(err)
}
```

See the `examples` folder for more examples on how to use the SDK.
