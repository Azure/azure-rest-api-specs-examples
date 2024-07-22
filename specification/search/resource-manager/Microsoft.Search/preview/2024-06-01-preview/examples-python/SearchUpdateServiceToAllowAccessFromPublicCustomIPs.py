from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.search import SearchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-search
# USAGE
    python search_update_service_to_allow_access_from_public_custom_ips.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SearchManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.services.update(
        resource_group_name="rg1",
        search_service_name="mysearchservice",
        service={
            "properties": {
                "networkRuleSet": {"ipRules": [{"value": "123.4.5.6"}, {"value": "123.4.6.0/18"}]},
                "partitionCount": 1,
                "publicNetworkAccess": "enabled",
                "replicaCount": 3,
            }
        },
    )
    print(response)


# x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/SearchUpdateServiceToAllowAccessFromPublicCustomIPs.json
if __name__ == "__main__":
    main()
