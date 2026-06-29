const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: 2026-03-15-preview/Job/CommandJob/createOrUpdate.json
 */
async function createOrUpdateCommandJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("test-rg", "my-aml-workspace", "string", {
    properties: {
      description: "string",
      codeId: "string",
      command: "string",
      computeId: "string",
      displayName: "string",
      distribution: { distributionType: "TensorFlow", parameterServerCount: 1, workerCount: 1 },
      environmentId: "string",
      environmentVariables: { string: "string" },
      experimentName: "string",
      identity: { identityType: "AMLToken" },
      inputs: { string: { description: "string", jobInputType: "literal", value: "string" } },
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
      parentJobName: "ParentRun",
      properties: { string: "string" },
      resources: {
        instanceCount: 1,
        instanceType: "string",
        properties: { string: { "e6b6493e-7d5e-4db3-be1e-306ec641327e": null } },
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
  });
  console.log(result);
}
