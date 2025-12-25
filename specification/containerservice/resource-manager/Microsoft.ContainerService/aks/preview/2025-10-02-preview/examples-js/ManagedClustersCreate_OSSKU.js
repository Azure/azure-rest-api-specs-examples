const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a managed cluster.
 *
 * @summary Creates or updates a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/ManagedClustersCreate_OSSKU.json
 */
async function createManagedClusterWithOssku() {
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
        osSKU: "AzureLinux",
        osType: "Linux",
        vmSize: "Standard_DS2_v2",
      },
    ],
    autoScalerProfile: { scaleDownDelayAfterAdd: "15m", scanInterval: "20s" },
    diskEncryptionSetID:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des",
    dnsPrefix: "dnsprefix1",
    enableRbac: true,
    httpProxyConfig: {
      httpProxy: "http://myproxy.server.com:8080",
      httpsProxy: "https://myproxy.server.com:8080",
      noProxy: ["localhost", "127.0.0.1"],
      trustedCa: "Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U=",
    },
    kubernetesVersion: "",
    linuxProfile: {
      adminUsername: "azureuser",
      ssh: { publicKeys: [{ keyData: "keydata" }] },
    },
    location: "location1",
    networkProfile: {
      loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
      loadBalancerSku: "standard",
      outboundType: "loadBalancer",
    },
    servicePrincipalProfile: { clientId: "clientid", secret: "secret" },
    sku: { name: "Basic", tier: "Free" },
    tags: { archv2: "", tier: "production" },
    windowsProfile: {
      adminPassword: "replacePassword1234$",
      adminUsername: "azureuser",
    },
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
