const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: 2026-03-15-preview/Job/AutoMLJob/createOrUpdate.json
 */
async function createOrUpdateAutoMLJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("test-rg", "my-aml-workspace", "string", {
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
        properties: { string: { "9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad": null } },
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
  });
  console.log(result);
}
