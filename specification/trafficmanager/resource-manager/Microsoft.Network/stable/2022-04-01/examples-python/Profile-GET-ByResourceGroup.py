from azure.identity import DefaultAzureCredential
from azure.mgmt.trafficmanager import TrafficManagerManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-trafficmanager
# USAGE
    python profile_get_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TrafficManagerManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.profiles.list_by_resource_group(
        resource_group_name="azuresdkfornetautoresttrafficmanager3640",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-GET-ByResourceGroup.json
if __name__ == "__main__":
    main()
