from azure.identity import DefaultAzureCredential

from azure.mgmt.batch import BatchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-batch
# USAGE
    python certificate_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BatchManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.certificate.begin_delete(
        resource_group_name="default-azurebatch-japaneast",
        account_name="sampleacct",
        certificate_name="sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e",
    ).result()


# x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/CertificateDelete.json
if __name__ == "__main__":
    main()
