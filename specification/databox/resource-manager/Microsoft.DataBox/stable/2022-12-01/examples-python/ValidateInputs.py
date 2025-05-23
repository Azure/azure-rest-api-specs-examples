from azure.identity import DefaultAzureCredential

from azure.mgmt.databox import DataBoxManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databox
# USAGE
    python validate_inputs.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="YourSubscriptionId",
    )

    response = client.service.validate_inputs(
        location="westus",
        validation_request={
            "individualRequestDetails": [
                {
                    "dataImportDetails": [
                        {
                            "accountDetails": {
                                "dataAccountType": "StorageAccount",
                                "storageAccountId": "/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName",
                            }
                        }
                    ],
                    "deviceType": "DataBox",
                    "transferType": "ImportToAzure",
                    "validationType": "ValidateDataTransferDetails",
                },
                {
                    "deviceType": "DataBox",
                    "shippingAddress": {
                        "addressType": "Commercial",
                        "city": "XXXX XXXX",
                        "companyName": "XXXX XXXX",
                        "country": "XX",
                        "postalCode": "00000",
                        "stateOrProvince": "XX",
                        "streetAddress1": "XXXX XXXX",
                        "streetAddress2": "XXXX XXXX",
                    },
                    "transportPreferences": {"preferredShipmentType": "MicrosoftManaged"},
                    "validationType": "ValidateAddress",
                },
                {"validationType": "ValidateSubscriptionIsAllowedToCreateJob"},
                {
                    "country": "XX",
                    "deviceType": "DataBox",
                    "location": "westus",
                    "transferType": "ImportToAzure",
                    "validationType": "ValidateSkuAvailability",
                },
                {"deviceType": "DataBox", "validationType": "ValidateCreateOrderLimit"},
                {
                    "deviceType": "DataBox",
                    "preference": {"transportPreferences": {"preferredShipmentType": "MicrosoftManaged"}},
                    "validationType": "ValidatePreferences",
                },
            ],
            "validationCategory": "JobCreationValidation",
        },
    )
    print(response)


# x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/ValidateInputs.json
if __name__ == "__main__":
    main()
