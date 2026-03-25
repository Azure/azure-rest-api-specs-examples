const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_DualStackNetworking.json
 */
async function createPoolDualStackNetworking() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create("default-azurebatch", "exampleacc", "dualstackpool", {
    vmSize: "Standard_D4ds_v5",
    networkConfiguration: {
      publicIPAddressConfiguration: { ipFamilies: ["IPv4", "IPv6"] },
      endpointConfiguration: {
        inboundNatPools: [
          {
            backendPort: 22,
            frontendPortRangeStart: 40000,
            frontendPortRangeEnd: 40500,
            name: "sshpool",
            protocol: "TCP",
            networkSecurityGroupRules: [
              {
                access: "Allow",
                priority: 1000,
                sourceAddressPrefix: "*",
                sourcePortRanges: ["*"],
              },
            ],
          },
        ],
      },
    },
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        imageReference: { publisher: "Canonical", offer: "ubuntu-24_04-lts", sku: "server" },
        nodeAgentSkuId: "batch.node.ubuntu 24.04",
      },
    },
    scaleSettings: { fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 } },
  });
  console.log(result);
}
