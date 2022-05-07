Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storageimportexport_2.0.1/sdk/storageimportexport/arm-storageimportexport/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

async function createImportJob() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const jobName = "myJob";
  const resourceGroupName = "myResourceGroup";
  const body = {
    location: "West US",
    properties: {
      backupDriveManifest: true,
      diagnosticsPath: "waimportexport",
      driveList: [
        {
          bitLockerKey: "238810-662376-448998-450120-652806-203390-606320-483076",
          driveHeaderHash:
            "0:1048576:FB6B6ED500D49DA6E0D723C98D42C657F2881CC13357C28DCECA6A524F1292501571A321238540E621AB5BD9C9A32637615919A75593E6CB5C1515DAE341CABF;135266304:143360:C957A189AFC38C4E80731252301EB91427CE55E61448FA3C73C6FDDE70ABBC197947EC8D0249A2C639BB10B95957D5820A4BE8DFBBF76FFFA688AE5CE0D42EC3",
          driveId: "9CA995BB",
          manifestFile: "8a0c23f7-14b7-470a-9633-fcd46590a1bc.manifest",
          manifestHash: "4228EC5D8E048CB9B515338C789314BE8D0B2FDBC7C7A0308E1C826242CDE74E",
        },
      ],
      jobType: "Import",
      logLevel: "Verbose",
      returnAddress: {
        city: "Redmond",
        countryOrRegion: "USA",
        email: "Test@contoso.com",
        phone: "4250000000",
        postalCode: "98007",
        recipientName: "Test",
        stateOrProvince: "wa",
        streetAddress1: "Street1",
        streetAddress2: "street2",
      },
      returnShipping: { carrierAccountNumber: "989ffff", carrierName: "FedEx" },
      storageAccountId:
        "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const result = await client.jobs.create(jobName, resourceGroupName, body);
  console.log(result);
}

createImportJob().catch(console.error);
```
