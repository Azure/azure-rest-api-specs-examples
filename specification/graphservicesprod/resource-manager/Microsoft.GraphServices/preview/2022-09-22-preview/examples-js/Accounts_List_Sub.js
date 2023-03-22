const { GraphServices } = require("@azure/arm-graphservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of accounts belonging to a subscription.
 *
 * @summary Returns list of accounts belonging to a subscription.
 * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_List_Sub.json
 */
async function getListOfAccountsBySubscription() {
  const subscriptionId =
    process.env["GRAPHSERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new GraphServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accounts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
