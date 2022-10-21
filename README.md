# Azure SDK examples for azure-rest-api-specs

This repository is the collection of Azure SDK examples corresponding to REST API examples from [Azure REST API Specifications](https://github.com/Azure/azure-rest-api-specs).

Involved SDK repositories:
- [Azure SDK for Go](https://github.com/Azure/azure-sdk-for-go)
- [Azure SDK for Java](https://github.com/Azure/azure-sdk-for-java)
- [Azure SDK for JavaScript](https://github.com/Azure/azure-sdk-for-js)
- [Azure SDK for Python](https://github.com/Azure/azure-sdk-for-python)
- [Azure SDK for .Net](https://github.com/Azure/azure-sdk-for-net)

The repository is maintained by automated pipeline.

## Specification on filename

Mapping rule of filename from examples from azure-rest-api-specs to azure-rest-api-specs-examples.
1. Replace `/examples/` with language specific folder of `/examples-<language>/`.
2. - [Metadata](#metadata): The JSON file as metadata.
   - [Code snippet](#code-snippet): Replace `.json` with `.<language-ext>` in file extension.

Currently supported `language` includes:
- `go` for Go
- `java` for Java
- `js` for JavaScript
- `python` for Python
- `dotnet` for .NET

Pending:
- `ts` for typescript

### An instance of the mapping

For instance, JSON example in azure-rest-api-specs with filename
```
specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
```
would map to code snippet (language=java) with filename
```
specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples-java/DataFlows_Create.java
```
and metadata with filename
```
specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples-java/DataFlows_Create.json
```
in azure-rest-api-specs-examples.

## File content

### Code snippet

Code snippet is runnable code in SDK language, as the SDK example.

Currently supported `language-ext` includes:
- `go` for Go
- `java` for Java
- `js` for JavaScript
- `py` for Python
- `cs` for .NET

Pending:
- `ts` for typescript

### Metadata

Metadata is a JSON that contains information related to the code snippet.

----
| Property | Type | Required | Description
-|-|-|-
| sdkUrl | string | yes | The URL to SDK documentation. It usually provides information on package and authentication.
----

### Cause of missing SDK examples

It is possible that for some api-version in some resource provider, there is no corresponding SDK example file, for some language.
Typical reason is that either the SDK of that language is not released, or the released SDK does not generate examples.

There is also possibility of missing a few SDK examples in a release.
Typical reason could be that SDK specifically removed a portion of the APIs, or the JSON example is considered not correct (e.g. some required parameter/property is not provided in JSON).

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
