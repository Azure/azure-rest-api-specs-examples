const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get build service agent pool.
 *
 * @summary Get build service agent pool.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildServiceAgentPool_Get.json
 */
async function buildServiceAgentPoolGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const agentPoolName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceAgentPool.get(
    resourceGroupName,
    serviceName,
    buildServiceName,
    agentPoolName
  );
  console.log(result);
}

buildServiceAgentPoolGet().catch(console.error);
