const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the agent pool in the Hybrid AKS provisioned cluster
 *
 * @summary Updates the agent pool in the Hybrid AKS provisioned cluster
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/UpdateAgentPool.json
 */
async function updateAgentPool() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
  const resourceGroupName =
    process.env["HYBRIDCONTAINERSERVICE_RESOURCE_GROUP"] || "test-arcappliance-resgrp";
  const resourceName = "test-hybridakscluster";
  const agentPoolName = "test-hybridaksnodepool";
  const agentPool = { count: 3, location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPoolOperations.update(
    resourceGroupName,
    resourceName,
    agentPoolName,
    agentPool
  );
  console.log(result);
}
