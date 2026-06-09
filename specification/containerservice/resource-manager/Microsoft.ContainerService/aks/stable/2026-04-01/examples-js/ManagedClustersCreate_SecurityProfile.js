const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-04-01/ManagedClustersCreate_SecurityProfile.json
 */
async function createManagedClusterWithSecurityProfileConfigured() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    location: "location1",
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
    kubernetesVersion: "",
    linuxProfile: { adminUsername: "azureuser", ssh: { publicKeys: [{ keyData: "keydata" }] } },
    networkProfile: {
      loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
      loadBalancerSku: "standard",
      outboundType: "loadBalancer",
    },
    securityProfile: {
      defender: {
        logAnalyticsWorkspaceResourceId:
          "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/microsoft.operationalinsights/workspaces/WORKSPACE_NAME",
        securityMonitoring: { enabled: true },
      },
      workloadIdentity: { enabled: true },
    },
    sku: { name: "Basic", tier: "Free" },
    tags: { archv2: "", tier: "production" },
  });
  console.log(result);
}
