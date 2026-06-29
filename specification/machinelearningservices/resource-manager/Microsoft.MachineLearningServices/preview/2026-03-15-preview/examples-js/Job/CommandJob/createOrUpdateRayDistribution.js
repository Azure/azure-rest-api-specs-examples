const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: 2026-03-15-preview/Job/CommandJob/createOrUpdateRayDistribution.json
 */
async function createOrUpdateCommandJobWithRayDistribution() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("test-rg", "my-aml-workspace", "string", {
    properties: {
      description: "Ray distributed command job",
      codeId: "string",
      command: "python train.py",
      computeId: "string",
      displayName: "ray-multi-node-job",
      distribution: {
        distributionType: "Ray",
        port: 6379,
        includeDashboard: true,
        dashboardPort: 8265,
      },
      environmentId: "string",
      environmentVariables: { string: "string" },
      experimentName: "ray-experiment",
      identity: { identityType: "AMLToken" },
      jobType: "Command",
      resources: { instanceCount: 2, instanceType: "string", properties: {} },
      tags: { string: "string" },
    },
  });
  console.log(result);
}
