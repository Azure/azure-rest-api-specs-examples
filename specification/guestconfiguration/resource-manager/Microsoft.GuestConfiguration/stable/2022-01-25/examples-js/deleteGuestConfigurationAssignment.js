const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a guest configuration assignment
 *
 * @summary Delete a guest configuration assignment
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/deleteGuestConfigurationAssignment.json
 */
async function deleteAnGuestConfigurationAssignment() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionId";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const guestConfigurationAssignmentName = "SecureProtocol";
  const vmName = "myVMName";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const result = await client.guestConfigurationAssignments.delete(
    resourceGroupName,
    guestConfigurationAssignmentName,
    vmName,
  );
  console.log(result);
}
