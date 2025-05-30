const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates an existing virtual endpoint. The request body can contain one to many of the properties present in the normal virtual endpoint definition.
 *
 * @summary Updates an existing virtual endpoint. The request body can contain one to many of the properties present in the normal virtual endpoint definition.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/VirtualEndpointUpdate.json
 */
async function updateAVirtualEndpointForAServerToUpdateThe() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "pgtestsvc4";
  const virtualEndpointName = "pgVirtualEndpoint1";
  const parameters = {
    endpointType: "ReadWrite",
    members: ["testReplica1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.virtualEndpoints.beginUpdateAndWait(
    resourceGroupName,
    serverName,
    virtualEndpointName,
    parameters,
  );
  console.log(result);
}
