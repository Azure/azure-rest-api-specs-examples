const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container.
 *
 * @summary Create or update container.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/DataContainer/createOrUpdate.json
 */
async function createOrUpdateWorkspaceDataContainer() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "testrg123";
  const workspaceName = "workspace123";
  const name = "datacontainer123";
  const body = {
    properties: {
      description: "string",
      dataType: "UriFile",
      properties: { properties1: "value1", properties2: "value2" },
      tags: { tag1: "value1", tag2: "value2" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.dataContainers.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    body,
  );
  console.log(result);
}
