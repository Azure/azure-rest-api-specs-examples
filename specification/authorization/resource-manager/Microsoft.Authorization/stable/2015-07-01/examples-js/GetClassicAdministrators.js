const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets service administrator, account administrator, and co-administrators for the subscription.
 *
 * @summary Gets service administrator, account administrator, and co-administrators for the subscription.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetClassicAdministrators.json
 */
async function listClassicAdministrators() {
  const subscriptionId = process.env["AUTHORIZATION_SUBSCRIPTION_ID"] || "subId";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.classicAdministrators.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
