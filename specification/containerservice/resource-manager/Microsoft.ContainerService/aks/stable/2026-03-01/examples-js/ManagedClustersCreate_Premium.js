const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-03-01/ManagedClustersCreate_Premium.json
 */
async function createManagedClusterWithLongTermSupport() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    location: "location1",
    addonProfiles: {},
    agentPoolProfiles: [
      {
        name: "nodepool1",
        type: "VirtualMachineScaleSets",
        count: 3,
        enableEncryptionAtHost: true,
        enableNodePublicIP: true,
        mode: "System",
        osType: "Linux",
        vmSize: "Standard_DS2_v2",
      },
    ],
    apiServerAccessProfile: { disableRunCommand: true },
    autoScalerProfile: { scaleDownDelayAfterAdd: "15m", scanInterval: "20s" },
    dnsPrefix: "dnsprefix1",
    enableRbac: true,
    kubernetesVersion: "",
    linuxProfile: { adminUsername: "azureuser", ssh: { publicKeys: [{ keyData: "keydata" }] } },
    networkProfile: {
      loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
      loadBalancerSku: "standard",
      outboundType: "loadBalancer",
    },
    servicePrincipalProfile: { clientId: "clientid", secret: "secret" },
    supportPlan: "AKSLongTermSupport",
    windowsProfile: { adminPassword: "replacePassword1234$", adminUsername: "azureuser" },
    sku: { name: "Base", tier: "Premium" },
    tags: { archv2: "", tier: "production" },
  });
  console.log(result);
}
