const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default,
  { getLongRunningPoller } = require("@azure-rest/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed cluster.
 *
 * @summary Creates or updates a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersCreate_HTTPProxy.json
 */
async function createManagedClusterWithHttpProxyConfigured() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = {
    body: {
      properties: {
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
        autoScalerProfile: { "scale-down-delay-after-add": "15m", "scan-interval": "20s" },
        diskEncryptionSetID:
          "/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des",
        dnsPrefix: "dnsprefix1",
        enablePodSecurityPolicy: true,
        enableRBAC: true,
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
        networkProfile: {
          loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
          loadBalancerSku: "standard",
          outboundType: "loadBalancer",
        },
        servicePrincipalProfile: { clientId: "clientid", secret: "secret" },
        windowsProfile: {
          adminPassword: "replacePassword1234$",
          adminUsername: "azureuser",
        },
      },
      location: "location1",

      sku: { name: "Basic", tier: "Free" },
      tags: { archv2: "", tier: "production" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = ContainerServiceManagementClient(credential);
  const initalResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}",
      subscriptionId,
      resourceGroupName,
      resourceName,
    )
    .put(parameters);
  const poller = await getLongRunningPoller(client, initalResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createManagedClusterWithHttpProxyConfigured().catch(console.error);
