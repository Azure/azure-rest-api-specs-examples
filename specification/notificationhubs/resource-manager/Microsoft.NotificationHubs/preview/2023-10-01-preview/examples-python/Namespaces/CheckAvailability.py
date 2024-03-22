from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.notificationhubs import NotificationHubsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-notificationhubs
# USAGE
    python check_availability.py

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

    response = client.namespaces.check_availability(
        parameters={"name": "sdk-Namespace-2924"},
    )
    print(response)


# x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/Namespaces/CheckAvailability.json
if __name__ == "__main__":
    main()
