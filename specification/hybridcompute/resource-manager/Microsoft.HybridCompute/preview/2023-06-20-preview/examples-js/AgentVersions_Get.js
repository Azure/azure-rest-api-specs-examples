const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Agent Versions along with the download link currently present.
 *
 * @summary Gets all Agent Versions along with the download link currently present.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/AgentVersions_Get.json
 */
async function getAgentVersions() {
  const osType = "myOsType";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential);
  const result = await client.agentVersionOperations.list(osType);
  console.log(result);
}
