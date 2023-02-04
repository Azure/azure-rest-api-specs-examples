const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an availability group listener.
 *
 * @summary Deletes an availability group listener.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/DeleteAvailabilityGroupListener.json
 */
async function deletesAnAvailabilityGroupListener() {
  const subscriptionId =
    process.env["SQLVIRTUALMACHINE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQLVIRTUALMACHINE_RESOURCE_GROUP"] || "testrg";
  const sqlVirtualMachineGroupName = "testvmgroup";
  const availabilityGroupListenerName = "agl-test";
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.availabilityGroupListeners.beginDeleteAndWait(
    resourceGroupName,
    sqlVirtualMachineGroupName,
    availabilityGroupListenerName
  );
  console.log(result);
}
