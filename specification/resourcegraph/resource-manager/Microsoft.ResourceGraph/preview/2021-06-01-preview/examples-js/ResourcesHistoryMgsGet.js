const { ResourceGraphClient } = require("@azure/arm-resourcegraph");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all snapshots of a resource for a given time interval.
 *
 * @summary List all snapshots of a resource for a given time interval.
 * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesHistoryMgsGet.json
 */
async function resourceHistoryManagementGroupScopeQuery() {
  const request = {
    managementGroups: ["e927f598-c1d4-4f72-8541-95d83a6a4ac8", "ProductionMG"],
    options: {
      interval: {
        end: new Date("2020-11-12T01:25:00.0000000Z"),
        start: new Date("2020-11-12T01:00:00.0000000Z"),
      },
    },
    query: "where name =~ 'cpu-utilization' | project id, name, properties",
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceGraphClient(credential);
  const result = await client.resourcesHistory(request);
  console.log(result);
}

resourceHistoryManagementGroupScopeQuery().catch(console.error);
