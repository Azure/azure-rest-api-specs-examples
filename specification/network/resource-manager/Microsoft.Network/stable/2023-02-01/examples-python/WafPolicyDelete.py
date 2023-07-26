from azure.identity import DefaultAzureCredential
from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python waf_policy_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.web_application_firewall_policies.begin_delete(
        resource_group_name="rg1",
        policy_name="Policy1",
    ).result()


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/WafPolicyDelete.json
if __name__ == "__main__":
    main()
