from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python exascale_db_storage_vaults_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OracleDatabaseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.exascale_db_storage_vaults.begin_create(
        resource_group_name="rgopenapi",
        exascale_db_storage_vault_name="vmClusterName",
        resource={
            "location": "ltguhzffucaytqg",
            "properties": {
                "additionalFlashCacheInPercent": 0,
                "description": "dmnvnnduldfmrmkkvvsdtuvmsmruxzzpsfdydgytlckutfozephjygjetrauvbdfcwmti",
                "displayName": "hbsybtelyvhpalemszcvartlhwvskrnpiveqfblvkdihoytqaotdgsgauvgivzaftfgeiwlyeqzssicwrrnlxtsmeakbcsxabjlt",
                "highCapacityDatabaseStorage": {"availableSizeInGbs": 28, "totalSizeInGbs": 16},
                "highCapacityDatabaseStorageInput": {"totalSizeInGbs": 21},
                "lifecycleState": "Provisioning",
                "ocid": "ocid1.autonomousdatabase.oc1..aaaaa3klq",
                "timeZone": "ltrbozwxjunncicrtzjrpqnqrcjgghohztrdlbfjrbkpenopyldwolslwgrgumjfkyovvkzcuxjujuxtjjzubvqvnhrswnbdgcbslopeofmtepbrrlymqwwszvsglmyuvlcuejshtpokirwklnwpcykhyinjmlqvxtyixlthtdishhmtipbygsayvgqzfrprgppylydlcskbmvwctxifdltippfvsxiughqbojqpqrekxsotnqsk",
            },
            "tags": {"key7827": "xqi"},
            "zones": ["qk"],
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-03-01/ExascaleDbStorageVaults_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
