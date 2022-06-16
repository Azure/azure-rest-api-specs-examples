const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the chap setting.
 *
 * @summary Creates or updates the chap setting.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ChapSettingsCreateOrUpdate.json
 */
async function chapSettingsCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-WSJQERQW3F";
  const chapUserName = "ChapSettingForSDK";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const chapSetting = {
    name: "ChapSettingForSDK",
    password: {
      encryptionAlgorithm: "RSAES_PKCS1_v_1_5",
      encryptionCertificateThumbprint: "D73DB57C4CDD6761E159F8D1E8A7D759424983FD",
      value: "<value>",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.chapSettingsOperations.beginCreateOrUpdateAndWait(
    deviceName,
    chapUserName,
    resourceGroupName,
    managerName,
    chapSetting
  );
  console.log(result);
}

chapSettingsCreateOrUpdate().catch(console.error);
