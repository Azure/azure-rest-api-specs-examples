from azure.identity import DefaultAzureCredential
from azure.mgmt.notificationhubs import NotificationHubsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-notificationhubs
# USAGE
    python name_space_list_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NotificationHubsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="29cfa613-cbbc-4512-b1d6-1b3a92c7fa40",
    )

    response = client.namespaces.list(
        resource_group_name="5ktrial",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceListByResourceGroup.json
if __name__ == "__main__":
    main()
