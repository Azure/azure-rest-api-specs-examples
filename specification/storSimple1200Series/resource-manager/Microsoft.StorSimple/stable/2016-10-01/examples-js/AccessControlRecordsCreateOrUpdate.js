const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates an access control record.
 *
 * @summary Creates or Updates an access control record.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/AccessControlRecordsCreateOrUpdate.json
 */
async function accessControlRecordsCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const accessControlRecordName = "AcrForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const accessControlRecord = {
    name: "AcrForSDKTest",
    initiatorName: "iqn.2017-06.com.contoso:ForTest",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.accessControlRecords.beginCreateOrUpdateAndWait(
    accessControlRecordName,
    resourceGroupName,
    managerName,
    accessControlRecord
  );
  console.log(result);
}

accessControlRecordsCreateOrUpdate().catch(console.error);
