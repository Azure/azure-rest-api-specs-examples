from azure.identity import DefaultAzureCredential

from azure.mgmt.elasticsan import ElasticSanMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elasticsan
# USAGE
    python elastic_sans_update_maximum_set_gen.py

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

    response = client.elastic_sans.begin_update(
        resource_group_name="resourcegroupname",
        elastic_san_name="elasticsanname",
        parameters={
            "properties": {
                "autoScaleProperties": {
                    "scaleUpProperties": {
                        "autoScalePolicyEnforcement": "None",
                        "capacityUnitScaleUpLimitTiB": 17,
                        "increaseCapacityUnitByTiB": 4,
                        "unusedSizeTiB": 24,
                    }
                },
                "baseSizeTiB": 13,
                "extendedCapacitySizeTiB": 29,
                "publicNetworkAccess": "Enabled",
            },
            "tags": {"key1931": "yhjwkgmrrwrcoxblgwgzjqusch"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/ElasticSans_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
