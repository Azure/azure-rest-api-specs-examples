from azure.identity import DefaultAzureCredential
from azure.mgmt.communication import CommunicationServiceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-communication
# USAGE
    python update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CommunicationServiceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345",
    )

    response = client.communication_services.begin_update(
        resource_group_name="MyResourceGroup",
        communication_service_name="MyCommunicationResource",
        parameters={"tags": {"newTag": "newVal"}},
    ).result()
    print(response)


# x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/update.json
if __name__ == "__main__":
    main()
