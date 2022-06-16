const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountEnableAd() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9407";
  const accountName = "sto8596";
  const parameters = {
    azureFilesIdentityBasedAuthentication: {
      activeDirectoryProperties: {
        accountType: "User",
        azureStorageSid: "S-1-5-21-2400535526-2334094090-2402026252-0012",
        domainGuid: "aebfc118-9fa9-4732-a21f-d98e41a77ae1",
        domainName: "adtest.com",
        domainSid: "S-1-5-21-2400535526-2334094090-2402026252",
        forestName: "adtest.com",
        netBiosDomainName: "adtest.com",
        samAccountName: "sam12498",
      },
      directoryServiceOptions: "AD",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.update(resourceGroupName, accountName, parameters);
  console.log(result);
}

storageAccountEnableAd().catch(console.error);
