const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all reports for the guest configuration assignment, latest report first.
 *
 * @summary List all reports for the guest configuration assignment, latest report first.
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listAllGuestConfigurationConnectedVMwarevSphereAssignmentsReports.json
 */
async function listAllGuestConfigurationAssignmentsForAVirtualMachine() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionid";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const vmName = "myVMName";
  const guestConfigurationAssignmentName = "AuditSecureProtocol";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const result = await client.guestConfigurationConnectedVMwarevSphereAssignmentsReports.list(
    resourceGroupName,
    vmName,
    guestConfigurationAssignmentName,
  );
  console.log(result);
}
