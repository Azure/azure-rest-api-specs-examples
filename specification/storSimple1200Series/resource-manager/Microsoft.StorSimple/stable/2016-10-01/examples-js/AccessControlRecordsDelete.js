const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the access control record.
 *
 * @summary Deletes the access control record.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/AccessControlRecordsDelete.json
 */
async function accessControlRecordsDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const accessControlRecordName = "AcrForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.accessControlRecords.beginDeleteAndWait(
    accessControlRecordName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

accessControlRecordsDelete().catch(console.error);
