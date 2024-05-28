from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python bastion_shareable_link_delete_by_token.py

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

    client.begin_delete_bastion_shareable_link_by_token(
        resource_group_name="rg1",
        bastion_host_name="bastionhosttenant",
        bsl_token_request={"tokens": ["abcd1234-efgh-hijk-5678-abcdefgh1234", "dcba4321-hgfe-kjih-8765-hgfedcba4321"]},
    ).result()


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/BastionShareableLinkDeleteByToken.json
if __name__ == "__main__":
    main()
