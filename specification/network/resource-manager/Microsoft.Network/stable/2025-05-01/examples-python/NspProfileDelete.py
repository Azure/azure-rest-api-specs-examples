from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python nsp_profile_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId",
    )

    client.network_security_perimeter_profiles.delete(
        resource_group_name="rg1",
        network_security_perimeter_name="nsp1",
        profile_name="profile1",
    )


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspProfileDelete.json
if __name__ == "__main__":
    main()
