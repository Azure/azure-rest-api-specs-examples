const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists Jobs in the workspace.
 *
 * @summary lists Jobs in the workspace.
 * x-ms-original-file: 2026-03-15-preview/Job/AutoMLJob/list.json
 */
async function listAutoMLJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.jobs.list("test-rg", "my-aml-workspace")) {
    resArray.push(item);
  }

  console.log(resArray);
}
