Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List accounts in ResourceGroup
 *
 * @summary List accounts in ResourceGroup
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListByResourceGroup.json
 */
async function accountsListByResourceGroup() {
async function accountsListByResourceGroup() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "SampleResourceGroup";
  const resourceGroupName = "SampleResourceGroup";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const client = new PurviewManagementClient(credential, subscriptionId);
  const resArray = new Array();
  const resArray = new Array();
  for await (let item of client.accounts.listByResourceGroup(resourceGroupName)) {
  for await (let item of client.accounts.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
    resArray.push(item);
  }
  }
  console.log(resArray);
  console.log(resArray);
}
}

accountsListByResourceGroup().catch(console.error);
accountsListByResourceGroup().catch(console.error);
```
