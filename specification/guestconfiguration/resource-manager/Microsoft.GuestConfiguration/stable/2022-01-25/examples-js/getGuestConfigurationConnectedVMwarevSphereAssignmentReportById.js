const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a report for the guest configuration assignment, by reportId.
 *
 * @summary Get a report for the guest configuration assignment, by reportId.
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/getGuestConfigurationConnectedVMwarevSphereAssignmentReportById.json
 */
async function getAGuestConfigurationAssignmentReportByIdForAVirtualMachine() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionid";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const vmName = "myvm";
  const guestConfigurationAssignmentName = "AuditSecureProtocol";
  const reportId = "7367cbb8-ae99-47d0-a33b-a283564d2cb1";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const result = await client.guestConfigurationConnectedVMwarevSphereAssignmentsReports.get(
    resourceGroupName,
    vmName,
    guestConfigurationAssignmentName,
    reportId,
  );
  console.log(result);
}
