const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the access control record.
 *
 * @summary Deletes the access control record.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/AccessControlRecordsDelete.json
 */
async function accessControlRecordsDelete() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const accessControlRecordName = "ACRForTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.accessControlRecords.beginDeleteAndWait(
    accessControlRecordName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

accessControlRecordsDelete().catch(console.error);
