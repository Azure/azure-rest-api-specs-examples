from azure.identity import DefaultAzureCredential

from azure.mgmt.privatedns import PrivateDnsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-privatedns
# USAGE
    python private_zone_patch.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PrivateDnsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.private_zones.begin_update(
        resource_group_name="resourceGroup1",
        private_zone_name="privatezone1.com",
        parameters={"tags": {"key2": "value2"}},
    ).result()
    print(response)


# x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/PrivateZonePatch.json
if __name__ == "__main__":
    main()
