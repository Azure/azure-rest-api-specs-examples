const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 *
 * @summary patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 * x-ms-original-file: 2026-05-01-preview/Clusters_Patch_AggregatorOrSingleRackDefinition.json
 */
async function patchClusterAggregatorOrSingleRackDefinition() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.clusters.update("resourceGroupName", "clusterName", {
    clusterUpdateParameters: {
      aggregatorOrSingleRackDefinition: {
        bareMetalMachineConfigurationData: [
          {
            bmcCredentials: { password: "{password}", username: "username" },
            bmcMacAddress: "AA:BB:CC:DD:EE:FF",
            bootMacAddress: "00:BB:CC:DD:EE:FF",
            machineDetails: "extraDetails",
            machineName: "bmmName1",
            rackSlot: 1,
            serialNumber: "BM1219XXX",
          },
          {
            bmcCredentials: { password: "{password}", username: "username" },
            bmcMacAddress: "AA:BB:CC:DD:EE:00",
            bootMacAddress: "00:BB:CC:DD:EE:00",
            machineDetails: "extraDetails",
            machineName: "bmmName2",
            rackSlot: 2,
            serialNumber: "BM1219YYY",
          },
        ],
        networkRackId:
          "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName",
        rackLocation: "Foo Datacenter, Floor 3, Aisle 9, Rack 2",
        rackSerialNumber: "newSerialNumber",
        rackSkuId:
          "/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName",
        storageApplianceConfigurationData: [
          {
            adminCredentials: { password: "{password}", username: "username" },
            rackSlot: 1,
            serialNumber: "BM1219XXX",
            storageApplianceName: "vmName",
          },
        ],
      },
      computeDeploymentThreshold: { grouping: "PerCluster", type: "PercentSuccess", value: 90 },
      tags: { key1: "myvalue1", key2: "myvalue2" },
    },
  });
  console.log(result);
}
