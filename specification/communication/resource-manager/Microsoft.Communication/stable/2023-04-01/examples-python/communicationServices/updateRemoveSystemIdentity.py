from azure.identity import DefaultAzureCredential
from azure.mgmt.communication import CommunicationServiceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-communication
# USAGE
    python update_remove_system_identity.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CommunicationServiceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="11112222-3333-4444-5555-666677778888",
    )

    response = client.communication_services.update(
        resource_group_name="MyResourceGroup",
        communication_service_name="MyCommunicationResource",
        parameters={"identity": {"type": "None"}, "tags": {"newTag": "newVal"}},
    )
    print(response)


# x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/communicationServices/updateRemoveSystemIdentity.json
if __name__ == "__main__":
    main()
