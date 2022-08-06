const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a configuration profile assignment
 *
 * @summary Delete a configuration profile assignment
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileAssignment.json
 */
async function deleteAnConfigurationProfileAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const configurationProfileAssignmentName = "default";
  const vmName = "myVMName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfileAssignments.delete(
    resourceGroupName,
    configurationProfileAssignmentName,
    vmName
  );
  console.log(result);
}

deleteAnConfigurationProfileAssignment().catch(console.error);
