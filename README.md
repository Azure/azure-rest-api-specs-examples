# Azure SDK examples for azure-rest-api-specs

This repository is the collection of Azure SDK examples corresponding to REST API examples from [Azure REST API Specifications](https://github.com/Azure/azure-rest-api-specs).

Involved SDK resositories:
- [Azure SDK for Go](https://github.com/Azure/azure-sdk-for-go)
- [Azure SDK for Java](https://github.com/Azure/azure-sdk-for-java)
- [Azure SDK for JavaScript](https://github.com/Azure/azure-sdk-for-js)

The repository is maintained by automated pipeline.

## Specification on filename

Mapping rule of filename from examples from azure-rest-api-specs to azure-rest-api-specs-examples.
1. Replace `.json` with `.md` in file extension.
2. Replace `/examples/` with language specific folder of `/examples-<language>/`.

Currently supported `language` includes:
* `go` for Go
* `java` for Java
* `js` for JavaScript

Pending:
* `python` for Python
* `dotnet` for .NET

For example, filename
```
specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
```
in azure-rest-api-specs would map to filename (language=java) 
```
specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples-java/DataFlows_Create.md
```
in azure-rest-api-specs-examples.

It is possible that for some api-version in some resource provider, there is no corresponding SDK example file, for some language.
Typical reason is that either the SDK of that language is not released, or the released SDK does not generate examples.

There is also possibility of missing a few SDK examples in a release.
Typical reason could be that SDK specifically removed a portion of the APIs, or the JSON example is considered not correct (e.g. some required parameter/property is not provided in JSON).

## Specification on file content

The file is markdown, for flexibility.

The typical content of the file includes 2 sections.

First section is the reference to SDK readme, for example
```
Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.7.0/sdk/resourcemanager/README.md) on how to add the SDK to your project and authenticate.
```

Generally SDK readme contains guide on how to get it from package manager, how to configure it in project, and how to authenticate with Azure.

Second section is a code block with syntax highlighting. This is the code of the SDK example.

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
