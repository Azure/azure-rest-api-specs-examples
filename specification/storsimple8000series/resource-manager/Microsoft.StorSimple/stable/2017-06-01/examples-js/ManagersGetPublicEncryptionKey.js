const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the symmetric encrypted public encryption key of the manager.
 *
 * @summary Returns the symmetric encrypted public encryption key of the manager.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersGetPublicEncryptionKey.json
 */
async function managersGetPublicEncryptionKey() {
  const subscriptionId = "d3ebfe71-b7a9-4c57-92b9-68a2afde4de5";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.managers.getPublicEncryptionKey(resourceGroupName, managerName);
  console.log(result);
}

managersGetPublicEncryptionKey().catch(console.error);
