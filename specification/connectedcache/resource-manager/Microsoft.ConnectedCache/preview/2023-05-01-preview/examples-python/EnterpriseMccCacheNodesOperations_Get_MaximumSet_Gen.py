from azure.identity import DefaultAzureCredential

from azure.mgmt.connectedcache import ConnectedCacheMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-connectedcache
# USAGE
    python enterprise_mcc_cache_nodes_operations_get_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConnectedCacheMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.enterprise_mcc_cache_nodes_operations.get(
        resource_group_name="rgConnectedCache",
        customer_resource_name="cygqjgtcetihsajgyqwwrbclssqsvhgltrboemcqqtpoxdbhykqxblaihmrumyhbsx",
        cache_node_resource_name="fqxfadsultwjfzdwlqkvneakfsbyhratytmssmiukpbnus",
    )
    print(response)


# x-ms-original-file: 2023-05-01-preview/EnterpriseMccCacheNodesOperations_Get_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
