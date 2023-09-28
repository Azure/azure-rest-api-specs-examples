const { PlaywrightTestingClient } = require("@azure/arm-playwrighttesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Account
 *
 * @summary Get a Account
 * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_Get.json
 */
async function accountsGet() {
  const subscriptionId =
    process.env["PLAYWRIGHTTESTING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["PLAYWRIGHTTESTING_RESOURCE_GROUP"] || "dummyrg";
  const name = "myPlaywrightAccount";
  const credential = new DefaultAzureCredential();
  const client = new PlaywrightTestingClient(credential, subscriptionId);
  const result = await client.accounts.get(resourceGroupName, name);
  console.log(result);
}
