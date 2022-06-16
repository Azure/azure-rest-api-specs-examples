const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an application package record and its associated binary file.
 *
 * @summary Deletes an application package record and its associated binary file.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationPackageDelete.json
 */
async function applicationPackageDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const applicationName = "app1";
  const versionName = "1";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.applicationPackageOperations.delete(
    resourceGroupName,
    accountName,
    applicationName,
    versionName
  );
  console.log(result);
}

applicationPackageDelete().catch(console.error);
