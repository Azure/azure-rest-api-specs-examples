const { PlaywrightTestingClient } = require("@azure/arm-playwrighttesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Account
 *
 * @summary Create a Account
 * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_CreateOrUpdate.json
 */
async function accountsCreateOrUpdate() {
  const subscriptionId =
    process.env["PLAYWRIGHTTESTING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["PLAYWRIGHTTESTING_RESOURCE_GROUP"] || "dummyrg";
  const name = "myPlaywrightAccount";
  const resource = {
    location: "westus",
    properties: { regionalAffinity: "Enabled" },
    tags: { team: "Dev Exp" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PlaywrightTestingClient(credential, subscriptionId);
  const result = await client.accounts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    resource,
  );
  console.log(result);
}
