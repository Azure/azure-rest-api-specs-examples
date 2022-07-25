const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This method does all necessary pre-job creation validation under subscription.
 *
 * @summary This method does all necessary pre-job creation validation under subscription.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateInputs.json
 */
async function validateInputs() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const location = "westus";
  const validationRequest = {
    individualRequestDetails: [
      {
        dataImportDetails: [
          {
            accountDetails: {
              dataAccountType: "StorageAccount",
              storageAccountId:
                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount",
            },
          },
        ],
        deviceType: "DataBox",
        transferType: "ImportToAzure",
        validationType: "ValidateDataTransferDetails",
      },
      {
        deviceType: "DataBox",
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
        transportPreferences: { preferredShipmentType: "MicrosoftManaged" },
        validationType: "ValidateAddress",
      },
      { validationType: "ValidateSubscriptionIsAllowedToCreateJob" },
      {
        country: "US",
        deviceType: "DataBox",
        location: "westus",
        transferType: "ImportToAzure",
        validationType: "ValidateSkuAvailability",
      },
      { deviceType: "DataBox", validationType: "ValidateCreateOrderLimit" },
      {
        deviceType: "DataBox",
        preference: {
          transportPreferences: { preferredShipmentType: "MicrosoftManaged" },
        },
        validationType: "ValidatePreferences",
      },
    ],
    validationCategory: "JobCreationValidation",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.service.validateInputs(location, validationRequest);
  console.log(result);
}

validateInputs().catch(console.error);
