const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Refreshes any information about the association.
 *
 * @summary Refreshes any information about the association.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/NetworkSecurityPerimeterConfigurationReconcile.json
 */
async function networkSecurityPerimeterConfigurationReconcile() {
  const subscriptionId =
    process.env["STORAGE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res4410";
  const accountName = "sto8607";
  const networkSecurityPerimeterConfigurationName =
    "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityPerimeterConfigurations.beginReconcileAndWait(
    resourceGroupName,
    accountName,
    networkSecurityPerimeterConfigurationName,
  );
  console.log(result);
}
