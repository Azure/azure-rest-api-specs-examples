const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This lists all the available Microsoft Support REST API operations.
 *
 * @summary This lists all the available Microsoft Support REST API operations.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListOperations.json
 */
async function getAllOperations() {
  const subscriptionId =
    process.env["SUPPORT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
