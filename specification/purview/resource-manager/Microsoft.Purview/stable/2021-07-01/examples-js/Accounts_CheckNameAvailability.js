const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if account name is available.
 *
 * @summary Checks if account name is available.
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_CheckNameAvailability.json
 */
async function accountsCheckNameAvailability() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const checkNameAvailabilityRequest = {
    name: "account1",
    type: "Microsoft.Purview/accounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.accounts.checkNameAvailability(checkNameAvailabilityRequest);
  console.log(result);
}

accountsCheckNameAvailability().catch(console.error);
