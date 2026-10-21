const { HorizonDbClient } = require("@azure/arm-horizondb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets information about a service-generated HorizonDB Cluster- or Pool-scoped Endpoint.
 *
 * @summary gets information about a service-generated HorizonDB Cluster- or Pool-scoped Endpoint.
 * x-ms-original-file: 2026-10-01-preview/Endpoints_GetReadOnly.json
 */
async function getAReadOnlyHorizonDBPoolEndpoint() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new HorizonDbClient(credential, subscriptionId);
  const result = await client.endpoints.get(
    "example-resource-group",
    "example-cluster",
    "example-pool-01",
  );
  console.log(result);
}
