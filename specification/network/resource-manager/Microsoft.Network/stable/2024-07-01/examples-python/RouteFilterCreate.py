from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python route_filter_create.py

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

    response = client.route_filters.begin_create_or_update(
        resource_group_name="rg1",
        route_filter_name="filterName",
        route_filter_parameters={
            "location": "West US",
            "properties": {
                "rules": [
                    {
                        "name": "ruleName",
                        "properties": {
                            "access": "Allow",
                            "communities": ["12076:5030", "12076:5040"],
                            "routeFilterRuleType": "Community",
                        },
                    }
                ]
            },
            "tags": {"key1": "value1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/RouteFilterCreate.json
if __name__ == "__main__":
    main()
