const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function devicesCreateOrUpdateSecuritySettings() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-T4ZA3EAJFR";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const securitySettings = {
    deviceAdminPassword: {
      encryptionAlgorithm: "RSAES_PKCS1_v_1_5",
      encryptionCertificateThumbprint: "D73DB57C4CDD6761E159F8D1E8A7D759424983FD",
      value: "<value>",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.devices.beginCreateOrUpdateSecuritySettingsAndWait(
    deviceName,
    resourceGroupName,
    managerName,
    securitySettings
  );
  console.log(result);
}

devicesCreateOrUpdateSecuritySettings().catch(console.error);
