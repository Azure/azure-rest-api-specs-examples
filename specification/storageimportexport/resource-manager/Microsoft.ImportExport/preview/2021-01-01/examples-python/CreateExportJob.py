from azure.identity import DefaultAzureCredential
from azure.mgmt.storageimportexport import StorageImportExport

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storageimportexport
# USAGE
    python create_export_job.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageImportExport(
        credential=DefaultAzureCredential(),
        subscription_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    )

    response = client.jobs.create(
        job_name="myExportJob",
        resource_group_name="myResourceGroup",
        body={
            "location": "West US",
            "properties": {
                "backupDriveManifest": True,
                "diagnosticsPath": "waimportexport",
                "export": {"blobList": {"blobPathPrefix": ["/"]}},
                "jobType": "Export",
                "logLevel": "Verbose",
                "returnAddress": {
                    "city": "Redmond",
                    "countryOrRegion": "USA",
                    "email": "Test@contoso.com",
                    "phone": "4250000000",
                    "postalCode": "98007",
                    "recipientName": "Test",
                    "stateOrProvince": "wa",
                    "streetAddress1": "Street1",
                    "streetAddress2": "street2",
                },
                "returnShipping": {"carrierAccountNumber": "989ffff", "carrierName": "FedEx"},
                "storageAccountId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/CreateExportJob.json
if __name__ == "__main__":
    main()
