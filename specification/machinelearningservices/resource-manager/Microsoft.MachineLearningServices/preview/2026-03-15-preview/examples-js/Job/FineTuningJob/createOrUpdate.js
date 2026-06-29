const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 *
 * @summary creates and executes a Job.
 * For update case, the Tags in the definition passed in will replace Tags in the existing job.
 * x-ms-original-file: 2026-03-15-preview/Job/FineTuningJob/createOrUpdate.json
 */
async function createOrUpdateFineTuningJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("test-rg", "my-aml-workspace", "string", {
    properties: {
      computeId: "gpu-compute",
      experimentName: "llm-finetuning",
      fineTuningDetails: {
        model: {
          description: undefined,
          jobInputType: "mlflow_model",
          mode: "ReadOnlyMount",
          uri: "azureml://registries/azureml-meta/models/Llama-2-7b/versions/11",
        },
        modelProvider: "Custom",
        taskType: "TextCompletion",
        trainingData: {
          description: undefined,
          jobInputType: "uri_file",
          mode: "ReadOnlyMount",
          uri: "azureml://registries/azureml-meta/models/Llama-2-7b/versions/11",
        },
      },
      jobType: "FineTuning",
      outputs: {
        string: {
          description: "string",
          jobOutputType: "mlflow_model",
          mode: "ReadWriteMount",
          uri: "string",
        },
      },
      queueSettings: { jobTier: "Standard" },
      resources: { instanceTypes: ["Standard_NC6"] },
    },
  });
  console.log(result);
}
