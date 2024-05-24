const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates and executes a Job.
 *
 * @summary Creates and executes a Job.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/CommandJob/createOrUpdate.json
 */
async function createOrUpdateCommandJob() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const id = "string";
  const body = {
    properties: {
      description: "string",
      codeId: "string",
      command: "string",
      computeId: "string",
      displayName: "string",
      distribution: {
        distributionType: "TensorFlow",
        parameterServerCount: 1,
        workerCount: 1,
      },
      environmentId: "string",
      environmentVariables: { string: "string" },
      experimentName: "string",
      identity: { identityType: "AMLToken" },
      inputs: {
        string: {
          description: "string",
          jobInputType: "literal",
          value: "string",
        },
      },
      jobType: "Command",
      limits: { jobLimitsType: "Command", timeout: "PT5M" },
      outputs: {
        string: {
          description: "string",
          jobOutputType: "uri_file",
          mode: "ReadWriteMount",
          uri: "string",
        },
      },
      properties: { string: "string" },
      resources: {
        instanceCount: 1,
        instanceType: "string",
        properties: {
          string: { "e6b6493e-7d5e-4db3-be1e-306ec641327e": null },
        },
      },
      services: {
        string: {
          endpoint: "string",
          jobServiceType: "string",
          port: 1,
          properties: { string: "string" },
        },
      },
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate(resourceGroupName, workspaceName, id, body);
  console.log(result);
}
