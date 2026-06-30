const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: 2026-03-15-preview/Job/PipelineJob/createOrUpdate.json
 */
async function createOrUpdatePipelineJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("test-rg", "my-aml-workspace", "string", {
    properties: {
      description: "string",
      computeId: "string",
      displayName: "string",
      experimentName: "string",
      inputs: { string: { description: "string", jobInputType: "literal", value: "string" } },
      jobType: "Pipeline",
      outputs: {
        string: { description: "string", jobOutputType: "uri_file", mode: "Upload", uri: "string" },
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
  });
  console.log(result);
}
