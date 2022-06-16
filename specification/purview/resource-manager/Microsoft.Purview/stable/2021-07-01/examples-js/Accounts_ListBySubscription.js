const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List accounts in Subscription
 *
 * @summary List accounts in Subscription
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListBySubscription.json
 */
async function accountsListBySubscription() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accounts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

accountsListBySubscription().catch(console.error);
