const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container.
 *
 * @summary Create or update container.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ModelContainer/createOrUpdate.json
 */
async function createOrUpdateModelContainer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspace123";
  const name = "testContainer";
  const body = {
    properties: {
      description: "Model container description",
      tags: { tag1: "value1", tag2: "value2" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.modelContainers.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    body
  );
  console.log(result);
}

createOrUpdateModelContainer().catch(console.error);
