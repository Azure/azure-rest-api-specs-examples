const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default,
  { getLongRunningPoller } = require("@azure-rest/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed cluster.
 *
 * @summary Creates or updates a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersCreate_DisableRunCommand.json
 */
async function createManagedClusterWithRunCommandDisabled() {
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
            enableEncryptionAtHost: true,
            enableNodePublicIP: true,
            mode: "System",
            osType: "Linux",
            vmSize: "Standard_DS2_v2",
          },
        ],
        apiServerAccessProfile: { disableRunCommand: true },
        autoScalerProfile: { "scale-down-delay-after-add": "15m", "scan-interval": "20s" },
        dnsPrefix: "dnsprefix1",
        enablePodSecurityPolicy: true,
        enableRBAC: true,
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
      resourceName
    )
    .put(parameters);
  const poller = getLongRunningPoller(client, initalResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createManagedClusterWithRunCommandDisabled().catch(console.error);
