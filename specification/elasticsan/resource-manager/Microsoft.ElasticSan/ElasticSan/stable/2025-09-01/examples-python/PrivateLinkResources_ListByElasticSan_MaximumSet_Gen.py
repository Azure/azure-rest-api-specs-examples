from azure.identity import DefaultAzureCredential

from azure.mgmt.elasticsan import ElasticSanMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elasticsan
# USAGE
    python private_link_resources_list_by_elastic_san_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ElasticSanMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.private_link_resources.list_by_elastic_san(
        resource_group_name="resourcegroupname",
        elastic_san_name="elasticsanname",
    )
    print(response)


# x-ms-original-file: 2025-09-01/PrivateLinkResources_ListByElasticSan_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
