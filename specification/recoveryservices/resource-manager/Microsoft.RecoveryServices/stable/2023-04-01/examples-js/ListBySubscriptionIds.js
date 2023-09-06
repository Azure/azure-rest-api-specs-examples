const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches all the resources of the specified type in the subscription.
 *
 * @summary Fetches all the resources of the specified type in the subscription.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ListBySubscriptionIds.json
 */
async function listOfRecoveryServicesResourcesInSubscriptionId() {
  const subscriptionId =
    process.env["RECOVERYSERVICES_SUBSCRIPTION_ID"] || "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vaults.listBySubscriptionId()) {
    resArray.push(item);
  }
  console.log(resArray);
}
