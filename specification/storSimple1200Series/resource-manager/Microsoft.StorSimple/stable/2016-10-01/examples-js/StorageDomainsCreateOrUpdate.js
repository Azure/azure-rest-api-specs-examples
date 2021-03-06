const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageDomainsCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const storageDomainName = "sd-fs-HSDK-4XY4FI2IVG";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const storageDomain = {
    name: "sd-fs-HSDK-4XY4FI2IVG",
    encryptionStatus: "Disabled",
    storageAccountCredentialIds: [
      "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageAccountCredentials/sacforsdktest",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.storageDomains.beginCreateOrUpdateAndWait(
    storageDomainName,
    resourceGroupName,
    managerName,
    storageDomain
  );
  console.log(result);
}

storageDomainsCreateOrUpdate().catch(console.error);
