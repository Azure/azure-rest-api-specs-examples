const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a configuration profile assignment
 *
 * @summary Get information about a configuration profile assignment
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileHCRPAssignment.json
 */
async function getAHcrpConfigurationProfileAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const machineName = "myMachineName";
  const configurationProfileAssignmentName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfileHcrpAssignments.get(
    resourceGroupName,
    machineName,
    configurationProfileAssignmentName
  );
  console.log(result);
}

getAHcrpConfigurationProfileAssignment().catch(console.error);
