const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-05-01/ManagedClustersCreate_Update.json
 */
async function createOrUpdateManagedCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "location1",
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
        scaleDownMode: "Deallocate",
        vmSize: "Standard_DS1_v2",
      },
    ],
    autoScalerProfile: {
      balanceSimilarNodeGroups: "true",
      expander: "priority",
      maxNodeProvisionTime: "15m",
      newPodScaleUpDelay: "1m",
      scaleDownDelayAfterAdd: "15m",
      scanInterval: "20s",
      skipNodesWithSystemPods: "false",
    },
    diskEncryptionSetID:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des",
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
    upgradeSettings: {
      overrideSettings: { forceUpgrade: false, until: new Date("2022-11-01T13:00:00Z") },
    },
    windowsProfile: { adminPassword: "replacePassword1234$", adminUsername: "azureuser" },
    sku: { name: "Basic", tier: "Free" },
    tags: { archv2: "", tier: "production" },
  });
  console.log(result);
}
