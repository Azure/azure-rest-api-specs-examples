const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to assess patches on a hybrid machine identity in Azure.
 *
 * @summary The operation to assess patches on a hybrid machine identity in Azure.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/machine/Machine_AssessPatches.json
 */
async function assessPatchStateOfAMachine() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroupName";
  const name = "myMachineName";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machines.beginAssessPatchesAndWait(resourceGroupName, name);
  console.log(result);
}
