const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a backup of the API Management service to the given Azure Storage Account. This is long running operation and could take several minutes to complete.
 *
 * @summary Creates a backup of the API Management service to the given Azure Storage Account. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackupWithUserAssignedManagedIdentity.json
 */
async function apiManagementBackupWithUserAssignedManagedIdentity() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const parameters = {
    accessType: "UserAssignedManagedIdentity",
    backupName: "backup5",
    clientId: "XXXXX-a154-4830-XXXX-46a12da1a1e2",
    containerName: "apim-backups",
    storageAccount: "contosorpstorage",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginBackupAndWait(
    resourceGroupName,
    serviceName,
    parameters
  );
  console.log(result);
}
