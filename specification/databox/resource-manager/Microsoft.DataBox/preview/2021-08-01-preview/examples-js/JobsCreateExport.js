const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new job with the specified parameters. Existing job cannot be updated with this API and should instead be updated with the Update job API.
 *
 * @summary Creates a new job with the specified parameters. Existing job cannot be updated with this API and should instead be updated with the Update job API.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsCreateExport.json
 */
async function jobsCreateExport() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "SdkRg8091";
  const jobName = "SdkJob6429";
  const jobResource = {
    location: "westus",
    sku: { name: "DataBox" },
    transferType: "ExportFromAzure",
    details: {
      contactDetails: {
        contactName: "Public SDK Test",
        emailList: ["testing@microsoft.com"],
        phone: "1234567890",
        phoneExtension: "1234",
      },
      dataExportDetails: [
        {
          accountDetails: {
            dataAccountType: "StorageAccount",
            storageAccountId:
              "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.Storage/storageAccounts/aaaaaa2",
          },
          transferConfiguration: {
            transferAllDetails: {
              include: {
                dataAccountType: "StorageAccount",
                transferAllBlobs: true,
                transferAllFiles: true,
              },
            },
            transferConfigurationType: "TransferAll",
          },
        },
      ],
      jobDetailsType: "DataBox",
      shippingAddress: {
        addressType: "Commercial",
        city: "San Francisco",
        companyName: "Microsoft",
        country: "US",
        postalCode: "94107",
        stateOrProvince: "CA",
        streetAddress1: "16 TOWNSEND ST",
        streetAddress2: "Unit 1",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.beginCreateAndWait(resourceGroupName, jobName, jobResource);
  console.log(result);
}

jobsCreateExport().catch(console.error);
