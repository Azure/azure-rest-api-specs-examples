Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get Diagnostics Categories
 *
 * @summary Description for Get Diagnostics Categories
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_ListSiteDiagnosticCategoriesSlot.json
 */
async function listAppSlotDiagnosticCategories() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const siteName = "SampleApp";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diagnostics.listSiteDiagnosticCategories(
    resourceGroupName,
    siteName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAppSlotDiagnosticCategories().catch(console.error);
```
