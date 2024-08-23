const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates and executes a Job.
For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary Creates and executes a Job.
For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Job/SweepJob/createOrUpdate.json
 */
async function createOrUpdateSweepJob() {
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
      earlyTermination: {
        delayEvaluation: 1,
        evaluationInterval: 1,
        policyType: "MedianStopping",
      },
      experimentName: "string",
      jobType: "Sweep",
      limits: {
        jobLimitsType: "Sweep",
        maxConcurrentTrials: 1,
        maxTotalTrials: 1,
        trialTimeout: "PT1S",
      },
      objective: { goal: "Minimize", primaryMetric: "string" },
      properties: { string: "string" },
      samplingAlgorithm: { samplingAlgorithmType: "Grid" },
      searchSpace: { string: {} },
      services: {
        string: {
          endpoint: "string",
          jobServiceType: "string",
          port: 1,
          properties: { string: "string" },
        },
      },
      tags: { string: "string" },
      trial: {
        codeId: "string",
        command: "string",
        distribution: { distributionType: "Mpi", processCountPerInstance: 1 },
        environmentId: "string",
        environmentVariables: { string: "string" },
        resources: {
          instanceCount: 1,
          instanceType: "string",
          properties: {
            string: { "e6b6493e-7d5e-4db3-be1e-306ec641327e": null },
          },
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate(resourceGroupName, workspaceName, id, body);
  console.log(result);
}
