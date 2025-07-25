const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 *
 * @summary Patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Clusters_Patch_AggregatorOrSingleRackDefinition.json
 */
async function patchClusterAggregatorOrSingleRackDefinition() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const clusterUpdateParameters = {
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
    computeDeploymentThreshold: {
      type: "PercentSuccess",
      grouping: "PerCluster",
      value: 90,
    },
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const options = { clusterUpdateParameters };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAndWait(resourceGroupName, clusterName, options);
  console.log(result);
}
