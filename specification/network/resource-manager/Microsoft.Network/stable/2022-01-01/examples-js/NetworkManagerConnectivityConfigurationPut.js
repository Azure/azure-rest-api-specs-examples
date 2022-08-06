const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates/Updates a new network manager connectivity configuration
 *
 * @summary Creates/Updates a new network manager connectivity configuration
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectivityConfigurationPut.json
 */
async function connectivityConfigurationsPut() {
  const subscriptionId = "subscriptionA";
  const resourceGroupName = "myResourceGroup";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestConnectivityConfig";
  const connectivityConfiguration = {
    description: "Sample Configuration",
    appliesToGroups: [
      {
        groupConnectivity: "None",
        isGlobal: "False",
        networkGroupId:
          "subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1",
        useHubGateway: "True",
      },
    ],
    connectivityTopology: "HubAndSpoke",
    deleteExistingPeering: "True",
    hubs: [
      {
        resourceId:
          "subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig",
        resourceType: "Microsoft.Network/virtualNetworks",
      },
    ],
    isGlobal: "True",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectivityConfigurations.createOrUpdate(
    resourceGroupName,
    networkManagerName,
    configurationName,
    connectivityConfiguration
  );
  console.log(result);
}

connectivityConfigurationsPut().catch(console.error);
