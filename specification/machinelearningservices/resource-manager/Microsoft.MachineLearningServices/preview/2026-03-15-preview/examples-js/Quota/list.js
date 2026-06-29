const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the currently assigned Workspace Quotas based on VMFamily.
 *
 * @summary gets the currently assigned Workspace Quotas based on VMFamily.
 * x-ms-original-file: 2026-03-15-preview/Quota/list.json
 */
async function listWorkspaceQuotasByVMFamily() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.quotas.list("eastus")) {
    resArray.push(item);
  }

  console.log(resArray);
}
