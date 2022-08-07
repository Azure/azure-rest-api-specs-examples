const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a configuration profile assignment
 *
 * @summary Delete a configuration profile assignment
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileHCRPAssignment.json
 */
async function deleteAHcrpConfigurationProfileAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const machineName = "myMachineName";
  const configurationProfileAssignmentName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfileHcrpAssignments.delete(
    resourceGroupName,
    machineName,
    configurationProfileAssignmentName
  );
  console.log(result);
}

deleteAHcrpConfigurationProfileAssignment().catch(console.error);
