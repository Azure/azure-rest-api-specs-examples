const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update version.
 *
 * @summary Create or update version.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/CodeVersion/createOrUpdate.json
 */
async function createOrUpdateWorkspaceCodeVersion() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const version = "string";
  const body = {
    properties: {
      description: "string",
      codeUri: "https://blobStorage/folderName",
      isAnonymous: false,
      properties: { string: "string" },
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.codeVersions.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    version,
    body,
  );
  console.log(result);
}
