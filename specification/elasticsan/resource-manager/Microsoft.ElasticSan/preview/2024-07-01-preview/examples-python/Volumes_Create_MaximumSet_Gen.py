from azure.identity import DefaultAzureCredential

from azure.mgmt.elasticsan import ElasticSanMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elasticsan
# USAGE
    python volumes_create_maximum_set_gen.py

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

    response = client.volumes.begin_create(
        resource_group_name="resourcegroupname",
        elastic_san_name="elasticsanname",
        volume_group_name="volumegroupname",
        volume_name="volumename",
        parameters={
            "properties": {
                "creationData": {"createSource": "None", "sourceId": "mdonegivjquite"},
                "managedBy": {"resourceId": "pclpkrpkpmvcsegcubrakcoodrubo"},
                "sizeGiB": 23,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/Volumes_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
