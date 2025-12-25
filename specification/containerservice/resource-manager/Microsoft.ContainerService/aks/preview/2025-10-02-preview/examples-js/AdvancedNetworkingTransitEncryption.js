const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a managed cluster.
 *
 * @summary Creates or updates a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/AdvancedNetworkingTransitEncryption.json
 */
async function createManagedClusterWithAdvancedNetworkingTransitEncryption() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const parameters = {
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
    dnsPrefix: "dnsprefix1",
    enableRbac: true,
    kubernetesVersion: "",
    linuxProfile: {
      adminUsername: "azureuser",
      ssh: { publicKeys: [{ keyData: "keydata" }] },
    },
    location: "location1",
    networkProfile: {
      advancedNetworking: {
        enabled: true,
        observability: { enabled: false },
        security: {
          advancedNetworkPolicies: "FQDN",
          enabled: true,
          transitEncryption: { type: "WireGuard" },
        },
      },
      loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
      loadBalancerSku: "standard",
      networkDataplane: "cilium",
      networkPlugin: "azure",
      networkPluginMode: "overlay",
      outboundType: "loadBalancer",
    },
    servicePrincipalProfile: { clientId: "clientid", secret: "secret" },
    sku: { name: "Basic", tier: "Free" },
    tags: { archv2: "", tier: "production" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    parameters,
  );
  console.log(result);
}
