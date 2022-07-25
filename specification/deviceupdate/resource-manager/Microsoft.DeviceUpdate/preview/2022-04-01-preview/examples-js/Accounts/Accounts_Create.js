const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates Account.
 *
 * @summary Creates or updates Account.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Accounts/Accounts_Create.json
 */
async function createsOrUpdatesAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "test-rg";
  const accountName = "contoso";
  const account = { location: "westus2" };
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.accounts.beginCreateAndWait(resourceGroupName, accountName, account);
  console.log(result);
}

createsOrUpdatesAccount().catch(console.error);
