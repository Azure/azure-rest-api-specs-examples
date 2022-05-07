Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storageimportexport_2.0.1/sdk/storageimportexport/arm-storageimportexport/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

async function createExportJob() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const jobName = "myExportJob";
  const resourceGroupName = "myResourceGroup";
  const body = {
    location: "West US",
    properties: {
      backupDriveManifest: true,
      diagnosticsPath: "waimportexport",
      export: { blobPathPrefix: ["/"] },
      jobType: "Export",
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

createExportJob().catch(console.error);
```
