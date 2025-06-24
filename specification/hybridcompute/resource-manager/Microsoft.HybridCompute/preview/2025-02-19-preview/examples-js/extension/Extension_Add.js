const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to Setup Machine Extensions.
 *
 * @summary The operation to Setup Machine Extensions.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/extension/Extension_Add.json
 */
async function setupMachineExtensions() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const machineName = "myMachine";
  const extensions = {
    extensions: [
      {
        type: "AzureMonitorAgentLinux",
        publisher: "Microsoft.Azure.Monitoring",
      },
      { type: "<extension_type>", publisher: "<extension_publisher>" },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.beginSetupExtensionsAndWait(
    resourceGroupName,
    machineName,
    extensions,
  );
  console.log(result);
}
