const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an EnvironmentVersion.
 *
 * @summary Creates or updates an EnvironmentVersion.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/EnvironmentVersion/createOrUpdate.json
 */
async function createOrUpdateEnvironmentVersion() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const version = "string";
  const body = {
    properties: {
      description: "string",
      build: {
        contextUri:
          "https://storage-account.blob.core.windows.net/azureml/DockerBuildContext/95ddede6b9b8c4e90472db3acd0a8d28/",
        dockerfilePath: "prod/Dockerfile",
      },
      condaFile: "string",
      image: "docker.io/tensorflow/serving:latest",
      inferenceConfig: {
        livenessRoute: { path: "string", port: 1 },
        readinessRoute: { path: "string", port: 1 },
        scoringRoute: { path: "string", port: 1 },
      },
      isAnonymous: false,
      properties: { string: "string" },
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.environmentVersions.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    version,
    body
  );
  console.log(result);
}

createOrUpdateEnvironmentVersion().catch(console.error);
