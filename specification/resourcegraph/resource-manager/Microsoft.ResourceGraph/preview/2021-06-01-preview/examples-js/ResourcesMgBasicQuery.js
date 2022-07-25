const { ResourceGraphClient } = require("@azure/arm-resourcegraph");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries the resources managed by Azure Resource Manager for scopes specified in the request.
 *
 * @summary Queries the resources managed by Azure Resource Manager for scopes specified in the request.
 * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesMgBasicQuery.json
 */
async function basicManagementGroupQuery() {
  const query = {
    managementGroups: ["e927f598-c1d4-4f72-8541-95d83a6a4ac8", "ProductionMG"],
    query: "Resources | project id, name, type, location, tags | limit 3",
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceGraphClient(credential);
  const result = await client.resources(query);
  console.log(result);
}

basicManagementGroupQuery().catch(console.error);
