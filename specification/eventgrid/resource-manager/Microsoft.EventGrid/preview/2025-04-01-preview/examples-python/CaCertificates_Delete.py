from azure.identity import DefaultAzureCredential

from azure.mgmt.eventgrid import EventGridManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventgrid
# USAGE
    python ca_certificates_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventGridManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="8f6b6269-84f2-4d09-9e31-1127efcd1e40",
    )

    client.ca_certificates.begin_delete(
        resource_group_name="examplerg",
        namespace_name="exampleNamespaceName1",
        ca_certificate_name="exampleCACertificateName1",
    ).result()


# x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/CaCertificates_Delete.json
if __name__ == "__main__":
    main()
