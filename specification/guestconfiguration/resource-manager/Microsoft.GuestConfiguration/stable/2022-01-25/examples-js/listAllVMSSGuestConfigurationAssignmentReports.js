const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all reports for the VMSS guest configuration assignment, latest report first.
 *
 * @summary List all reports for the VMSS guest configuration assignment, latest report first.
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listAllVMSSGuestConfigurationAssignmentReports.json
 */
async function listAllReportsForTheVmssGuestConfigurationAssignmentWithLatestReportFirst() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionid";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const vmssName = "myVMSSName";
  const name = "AuditSecureProtocol";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.guestConfigurationAssignmentReportsVmss.list(
    resourceGroupName,
    vmssName,
    name,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
