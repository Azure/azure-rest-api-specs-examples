from azure.identity import DefaultAzureCredential

from azure.mgmt.connectedcache import ConnectedCacheMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-connectedcache
# USAGE
    python isp_cache_nodes_operations_list_by_isp_customer_resource_maximum_set_gen.py

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

    response = client.isp_cache_nodes_operations.list_by_isp_customer_resource(
        resource_group_name="rgConnectedCache",
        customer_resource_name="MccRPTest1",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2023-05-01-preview/IspCacheNodesOperations_ListByIspCustomerResource_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
