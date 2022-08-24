const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes account.
 *
 * @summary Deletes account.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Accounts/Accounts_Delete.json
 */
async function deletesAnAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "test-rg";
  const accountName = "contoso";
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.accounts.beginDeleteAndWait(resourceGroupName, accountName);
  console.log(result);
}

deletesAnAccount().catch(console.error);
