const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Refreshes any information about the association.
 *
 * @summary Refreshes any information about the association.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/NetworkSecurityPerimeterConfigurationReconcile.json
 */
async function networkSecurityPerimeterConfigurationList() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "res4410";
  const accountName = "sto8607";
  const networkSecurityPerimeterConfigurationName =
    "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityPerimeterConfigurations.beginReconcileAndWait(
    resourceGroupName,
    accountName,
    networkSecurityPerimeterConfigurationName,
  );
  console.log(result);
}
