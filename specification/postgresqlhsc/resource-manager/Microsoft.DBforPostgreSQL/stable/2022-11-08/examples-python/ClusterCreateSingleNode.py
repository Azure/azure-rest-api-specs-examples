from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdbforpostgresql import CosmosdbForPostgresqlMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdbforpostgresql
# USAGE
    python cluster_create_single_node.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosdbForPostgresqlMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.clusters.begin_create(
        resource_group_name="TestGroup",
        cluster_name="testcluster-singlenode",
        parameters={
            "location": "westus",
            "properties": {
                "administratorLoginPassword": "password",
                "citusVersion": "11.3",
                "coordinatorEnablePublicIpAccess": True,
                "coordinatorServerEdition": "GeneralPurpose",
                "coordinatorStorageQuotaInMb": 131072,
                "coordinatorVCores": 8,
                "enableHa": True,
                "enableShardsOnCoordinator": True,
                "nodeCount": 0,
                "postgresqlVersion": "15",
                "preferredPrimaryZone": "1",
            },
            "tags": {"owner": "JohnDoe"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterCreateSingleNode.json
if __name__ == "__main__":
    main()
