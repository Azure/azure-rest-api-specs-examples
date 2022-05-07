Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets all static sites in the specified resource group.
 *
 * @summary Description for Gets all static sites in the specified resource group.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetStaticSites.json
 */
async function getStaticSitesForAResourceGroup() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.staticSites.listStaticSitesByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getStaticSitesForAResourceGroup().catch(console.error);
```
