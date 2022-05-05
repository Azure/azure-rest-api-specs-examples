Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of privately linkable resources for an account
 *
 * @summary Gets a list of privately linkable resources for an account
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_ListByAccount.json
 */
async function privateLinkResourcesListByAccount() {
async function privateLinkResourcesListByAccount() {
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const resourceGroupName = "SampleResourceGroup";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const accountName = "account1";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const client = new PurviewManagementClient(credential, subscriptionId);
  const resArray = new Array();
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listByAccount(
  for await (let item of client.privateLinkResources.listByAccount(
    resourceGroupName,
    resourceGroupName,
    accountName
    accountName
  )) {
  )) {
    resArray.push(item);
    resArray.push(item);
  }
  }
  console.log(resArray);
  console.log(resArray);
}
}

privateLinkResourcesListByAccount().catch(console.error);
privateLinkResourcesListByAccount().catch(console.error);
```
