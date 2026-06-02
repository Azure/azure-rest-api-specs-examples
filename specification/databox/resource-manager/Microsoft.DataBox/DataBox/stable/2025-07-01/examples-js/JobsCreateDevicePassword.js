const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new job with the specified parameters. Existing job cannot be updated with this API and should instead be updated with the Update job API.
 *
 * @summary creates a new job with the specified parameters. Existing job cannot be updated with this API and should instead be updated with the Update job API.
 * x-ms-original-file: 2025-07-01/JobsCreateDevicePassword.json
 */
async function jobsCreateDevicePassword() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "YourSubscriptionId";
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.create("YourResourceGroupName", "TestJobName1", {
    location: "westus",
    transferType: "ImportToAzure",
    details: {
      contactDetails: {
        contactName: "XXXX XXXX",
        emailList: ["xxxx@xxxx.xxx"],
        phone: "0000000000",
        phoneExtension: "",
      },
      dataImportDetails: [
        {
          accountDetails: {
            dataAccountType: "StorageAccount",
            sharePassword: "<sharePassword>",
            storageAccountId:
              "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName",
          },
        },
      ],
      devicePassword: "<devicePassword>",
      jobDetailsType: "DataBox",
      shippingAddress: {
        addressType: "Commercial",
        city: "XXXX XXXX",
        companyName: "XXXX XXXX",
        country: "XX",
        postalCode: "00000",
        stateOrProvince: "XX",
        streetAddress1: "XXXX XXXX",
        streetAddress2: "XXXX XXXX",
      },
    },
    sku: { name: "DataBox", model: "DataBox" },
  });
  console.log(result);
}
