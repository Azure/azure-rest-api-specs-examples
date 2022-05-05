Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an account
 *
 * @summary Creates or updates an account
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_CreateOrUpdate.json
 */
async function accountsCreateOrUpdate() {
async function accountsCreateOrUpdate() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "SampleResourceGroup";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const accountName = "account1";
  const account = {
  const account = {
    location: "West US 2",
    location: "West US 2",
    managedResourceGroupName: "custom-rgname",
    managedResourceGroupName: "custom-rgname",
  };
  };
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginCreateOrUpdateAndWait(
  const result = await client.accounts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceGroupName,
    accountName,
    accountName,
    account
    account
  );
  );
  console.log(result);
  console.log(result);
}
}

accountsCreateOrUpdate().catch(console.error);
accountsCreateOrUpdate().catch(console.error);
```
