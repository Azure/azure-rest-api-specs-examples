Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appcontainers_1.0.0/sdk/appcontainers/arm-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Container App SourceControls in a given resource group.
 *
 * @summary Get the Container App SourceControls in a given resource group.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/SourceControls_ListByContainer.json
 */
async function listAppSourceControls() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerAppsSourceControls.listByContainerApp(
    resourceGroupName,
    containerAppName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAppSourceControls().catch(console.error);
```
