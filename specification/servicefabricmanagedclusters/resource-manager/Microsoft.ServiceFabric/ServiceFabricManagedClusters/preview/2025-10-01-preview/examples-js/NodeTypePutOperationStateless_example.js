const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationStateless_example.json
 */
async function putAnStatelessNodeTypeWithTemporaryDiskForServiceFabric() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.createOrUpdate("resRg", "myCluster", "BE", {
    properties: {
      enableEncryptionAtHost: true,
      isPrimary: false,
      isStateless: true,
      multiplePlacementGroups: true,
      useTempDataDisk: true,
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
      vmInstanceCount: 10,
      vmSize: "Standard_DS3",
    },
  });
  console.log(result);
}
