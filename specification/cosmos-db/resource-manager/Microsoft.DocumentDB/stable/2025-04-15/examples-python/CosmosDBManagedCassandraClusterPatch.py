from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_managed_cassandra_cluster_patch.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.cassandra_clusters.begin_update(
        resource_group_name="cassandra-prod-rg",
        cluster_name="cassandra-prod",
        body={
            "properties": {
                "authenticationMethod": "None",
                "externalGossipCertificates": [
                    {"pem": "-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"}
                ],
                "externalSeedNodes": [
                    {"ipAddress": "10.52.221.2"},
                    {"ipAddress": "10.52.221.3"},
                    {"ipAddress": "10.52.221.4"},
                ],
                "hoursBetweenBackups": 12,
            },
            "tags": {"owner": "mike"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBManagedCassandraClusterPatch.json
if __name__ == "__main__":
    main()
