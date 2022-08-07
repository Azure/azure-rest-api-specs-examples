const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an association between a AzureStackHCI cluster and Automanage configuration profile
 *
 * @summary Creates an association between a AzureStackHCI cluster and Automanage configuration profile
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileHCIAssignment.json
 */
async function createOrUpdateAHciConfigurationProfileAssignment() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const clusterName = "myClusterName";
  const configurationProfileAssignmentName = "default";
  const parameters = {
    properties: {
      configurationProfile:
        "/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfileHCIAssignments.createOrUpdate(
    resourceGroupName,
    clusterName,
    configurationProfileAssignmentName,
    parameters
  );
  console.log(result);
}

createOrUpdateAHciConfigurationProfileAssignment().catch(console.error);
