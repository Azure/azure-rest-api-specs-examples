from azure.identity import DefaultAzureCredential

from azure.mgmt.elasticsan import ElasticSanMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elasticsan
# USAGE
    python private_endpoint_connections_list_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ElasticSanMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionid",
    )

    response = client.private_endpoint_connections.list(
        resource_group_name="resourcegroupname",
        elastic_san_name="elasticsanname",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/PrivateEndpointConnections_List_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
