const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates and executes a Job.
For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary Creates and executes a Job.
For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Job/AutoMLJob/createOrUpdate.json
 */
async function createOrUpdateAutoMlJob() {
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
      environmentId: "string",
      environmentVariables: { string: "string" },
      experimentName: "string",
      identity: { identityType: "AMLToken" },
      isArchived: false,
      jobType: "AutoML",
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
          string: { "9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad": null },
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
      taskDetails: {
        limitSettings: { maxTrials: 2 },
        modelSettings: { validationCropSize: 2 },
        searchSpace: [{ validationCropSize: "choice(2, 360)" }],
        targetColumnName: "string",
        taskType: "ImageClassification",
        trainingData: { jobInputType: "mltable", uri: "string" },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate(resourceGroupName, workspaceName, id, body);
  console.log(result);
}
