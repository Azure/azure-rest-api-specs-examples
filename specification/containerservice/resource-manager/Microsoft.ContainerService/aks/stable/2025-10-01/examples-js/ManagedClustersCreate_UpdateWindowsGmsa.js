const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a managed cluster.
 *
 * @summary Creates or updates a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/ManagedClustersCreate_UpdateWindowsGmsa.json
 */
async function createOrUpdateManagedClusterWithWindowsGMsaEnabled() {
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
        availabilityZones: ["1", "2", "3"],
        count: 3,
        enableNodePublicIP: true,
        mode: "System",
        osType: "Linux",
        vmSize: "Standard_DS1_v2",
      },
    ],
    autoScalerProfile: { scaleDownDelayAfterAdd: "15m", scanInterval: "20s" },
    diskEncryptionSetID:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des",
    dnsPrefix: "dnsprefix1",
    enableRbac: true,
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourcegroups/rgName1/providers/MicrosoftManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
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
      gmsaProfile: { enabled: true },
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
