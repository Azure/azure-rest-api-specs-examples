const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates and executes a Job.
 *
 * @summary Creates and executes a Job.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/PipelineJob/createOrUpdate.json
 */
async function createOrUpdatePipelineJob() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const id = "string";
  const body = {
    properties: {
      description: "string",
      computeId: "string",
      displayName: "string",
      experimentName: "string",
      inputs: {
        string: {
          description: "string",
          jobInputType: "literal",
          value: "string",
        },
      },
      jobType: "Pipeline",
      outputs: {
        string: {
          description: "string",
          jobOutputType: "uri_file",
          mode: "Upload",
          uri: "string",
        },
      },
      properties: { string: "string" },
      services: {
        string: {
          endpoint: "string",
          jobServiceType: "string",
          port: 1,
          properties: { string: "string" },
        },
      },
      settings: {},
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate(resourceGroupName, workspaceName, id, body);
  console.log(result);
}
