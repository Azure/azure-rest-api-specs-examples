Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an account
 *
 * @summary Updates an account
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Update.json
 */
async function accountsUpdate() {
async function accountsUpdate() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "SampleResourceGroup";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const accountName = "account1";
  const accountUpdateParameters = {
  const accountUpdateParameters = {
    tags: { newTag: "New tag value." },
    tags: { newTag: "New tag value." },
  };
  };
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginUpdateAndWait(
  const result = await client.accounts.beginUpdateAndWait(
    resourceGroupName,
    resourceGroupName,
    accountName,
    accountName,
    accountUpdateParameters
    accountUpdateParameters
  );
  );
  console.log(result);
  console.log(result);
}
}

accountsUpdate().catch(console.error);
accountsUpdate().catch(console.error);
```
