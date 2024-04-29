const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all guest configuration assignments for a subscription.
 *
 * @summary List all guest configuration assignments for a subscription.
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listSubGuestConfigurationAssignments.json
 */
async function listAllGuestConfigurationAssignmentsForASubscription() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.guestConfigurationAssignments.listSubscriptionList()) {
    resArray.push(item);
  }
  console.log(resArray);
}
