const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified storage account credential name.
 *
 * @summary Gets the properties of the specified storage account credential name.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/StorageAccountCredentialsGet.json
 */
async function storageAccountCredentialsGet() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const storageAccountCredentialName = "SACForTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.storageAccountCredentials.get(
    storageAccountCredentialName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

storageAccountCredentialsGet().catch(console.error);
