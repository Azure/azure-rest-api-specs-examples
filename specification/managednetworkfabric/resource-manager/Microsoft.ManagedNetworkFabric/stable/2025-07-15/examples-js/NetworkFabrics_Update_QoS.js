const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update certain properties of the Network Fabric resource.
 *
 * @summary update certain properties of the Network Fabric resource.
 * x-ms-original-file: 2025-07-15/NetworkFabrics_Update_QoS.json
 */
async function networkFabricsUpdateQoSEnable() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0000ABCD-0A0B-0000-0000-000000ABCDEF";
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkFabrics.update("example-rg", "example-fabric", {
    qosConfiguration: { qosConfigurationState: "Enabled" },
  });
  console.log(result);
}
