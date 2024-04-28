const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a guest configuration assignment
 *
 * @summary Get information about a guest configuration assignment
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/getGuestConfigurationConnectedVMwarevSphereAssignment.json
 */
async function getAGuestConfigurationAssignment() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionId";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const vmName = "myVMName";
  const guestConfigurationAssignmentName = "SecureProtocol";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const result = await client.guestConfigurationConnectedVMwarevSphereAssignments.get(
    resourceGroupName,
    vmName,
    guestConfigurationAssignmentName,
  );
  console.log(result);
}
