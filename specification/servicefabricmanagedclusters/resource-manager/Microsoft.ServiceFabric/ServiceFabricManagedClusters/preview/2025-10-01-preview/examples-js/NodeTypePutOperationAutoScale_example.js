const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationAutoScale_example.json
 */
async function putANodeTypeWithAutoScaleParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.createOrUpdate("resRg", "myCluster", "BE", {
    properties: {
      capacities: { ClientConnections: "65536" },
      dataDiskSizeGB: 200,
      dataDiskType: "Premium_LRS",
      isPrimary: false,
      isStateless: true,
      multiplePlacementGroups: true,
      placementProperties: {
        HasSSD: "true",
        NodeColor: "green",
        SomeProperty: "5",
      },
      vmExtensions: [
        {
          name: "Microsoft.Azure.Geneva.GenevaMonitoring",
          properties: {
            type: "GenevaMonitoring",
            autoUpgradeMinorVersion: true,
            publisher: "Microsoft.Azure.Geneva",
            settings: {},
            typeHandlerVersion: "2.0",
          },
        },
      ],
      vmImageOffer: "WindowsServer",
      vmImagePublisher: "MicrosoftWindowsServer",
      vmImageSku: "2016-Datacenter-Server-Core",
      vmImageVersion: "latest",
      vmInstanceCount: -1,
      vmManagedIdentity: {
        userAssignedIdentities: [
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2",
        ],
      },
      vmSecrets: [
        {
          sourceVault: {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault",
          },
          vaultCertificates: [
            {
              certificateStore: "My",
              certificateUrl:
                "https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c",
            },
          ],
        },
      ],
      vmSize: "Standard_DS3",
    },
  });
  console.log(result);
}
