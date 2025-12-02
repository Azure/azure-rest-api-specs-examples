const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationDedicatedHost_example.json
 */
async function putNodeTypeWithDedicatedHosts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.createOrUpdate("resRg", "myCluster", "BE", {
    properties: {
      capacities: {},
      dataDiskSizeGB: 200,
      dataDiskType: "StandardSSD_LRS",
      hostGroupId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testhostgroupRG/providers/Microsoft.Compute/hostGroups/testHostGroup",
      isPrimary: false,
      placementProperties: {},
      vmImageOffer: "WindowsServer",
      vmImagePublisher: "MicrosoftWindowsServer",
      vmImageSku: "2019-Datacenter",
      vmImageVersion: "latest",
      vmInstanceCount: 10,
      vmSize: "Standard_D8s_v3",
      zones: ["1"],
    },
  });
  console.log(result);
}
