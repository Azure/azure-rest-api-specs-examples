const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-08-02-preview/ManagedClustersCreate_EnableAgentSandbox.json
 */
async function createOrUpdateManagedClusterWithAgentSandboxEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    identity: { type: "SystemAssigned" },
    location: "location1",
    addonProfiles: {},
    agentPoolProfiles: [
      {
        name: "nodepool1",
        type: "VirtualMachineScaleSets",
        count: 3,
        enableNodePublicIP: true,
        mode: "System",
        osType: "Linux",
        vmSize: "Standard_DS2_v2",
      },
    ],
    agentSandboxProfile: { enabled: true },
    dnsPrefix: "dnsprefix1",
    enableRbac: true,
    kubernetesVersion: "1.35.6",
    linuxProfile: { adminUsername: "azureuser", ssh: { publicKeys: [{ keyData: "keydata" }] } },
    networkProfile: {
      loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
      loadBalancerSku: "standard",
      networkDataplane: "cilium",
      networkPlugin: "azure",
      networkPluginMode: "overlay",
      outboundType: "loadBalancer",
    },
    sku: { name: "Basic", tier: "Free" },
    tags: { archv2: "", tier: "production" },
  });
  console.log(result);
}
