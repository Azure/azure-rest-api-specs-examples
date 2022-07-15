const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update version.
 *
 * @summary Create or update version.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/ComponentVersion/createOrUpdate.json
 */
async function createOrUpdateComponentVersion() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const version = "string";
  const body = {
    properties: {
      description: "string",
      componentSpec: { "8ced901b-d826-477d-bfef-329da9672513": null },
      isAnonymous: false,
      properties: { string: "string" },
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.componentVersions.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    version,
    body
  );
  console.log(result);
}

createOrUpdateComponentVersion().catch(console.error);
