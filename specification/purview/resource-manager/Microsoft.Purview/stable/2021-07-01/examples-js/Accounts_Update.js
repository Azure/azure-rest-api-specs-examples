const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an account
 *
 * @summary Updates an account
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Update.json
 */
async function accountsUpdate() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const accountUpdateParameters = {
    tags: { newTag: "New tag value." },
  };
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    accountUpdateParameters
  );
  console.log(result);
}

accountsUpdate().catch(console.error);
