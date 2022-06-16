const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the configuration of a Service Fabric cluster resource with the specified name.
 *
 * @summary Update the configuration of a Service Fabric cluster resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPatchOperation_example.json
 */
async function patchACluster() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const parameters = {
    eventStoreServiceEnabled: true,
    nodeTypes: [
      {
        name: "nt1vm",
        applicationPorts: { endPort: 30000, startPort: 20000 },
        clientConnectionEndpointPort: 19000,
        durabilityLevel: "Bronze",
        ephemeralPorts: { endPort: 64000, startPort: 49000 },
        httpGatewayEndpointPort: 19007,
        isPrimary: true,
        vmInstanceCount: 5,
      },
      {
        name: "testnt1",
        applicationPorts: { endPort: 2000, startPort: 1000 },
        clientConnectionEndpointPort: 0,
        durabilityLevel: "Bronze",
        ephemeralPorts: { endPort: 4000, startPort: 3000 },
        httpGatewayEndpointPort: 0,
        isPrimary: false,
        vmInstanceCount: 3,
      },
    ],
    reliabilityLevel: "Bronze",
    tags: { a: "b" },
    upgradeMode: "Automatic",
    upgradePauseEndTimestampUtc: new Date("2021-06-25T22:00:00Z"),
    upgradePauseStartTimestampUtc: new Date("2021-06-21T22:00:00Z"),
    upgradeWave: "Wave",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

patchACluster().catch(console.error);
