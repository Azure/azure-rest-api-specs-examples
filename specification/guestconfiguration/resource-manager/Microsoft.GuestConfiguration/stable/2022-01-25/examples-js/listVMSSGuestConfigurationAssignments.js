const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all guest configuration assignments for VMSS.
 *
 * @summary List all guest configuration assignments for VMSS.
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listVMSSGuestConfigurationAssignments.json
 */
async function listAllGuestConfigurationAssignmentsForVmss() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionId";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const vmssName = "myVMSSName";
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.guestConfigurationAssignmentsVmss.list(
    resourceGroupName,
    vmssName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
