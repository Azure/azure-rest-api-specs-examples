from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_sql_container_create_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.sql_resources.begin_create_update_sql_container(
        resource_group_name="rg1",
        account_name="ddb1",
        database_name="databaseName",
        container_name="containerName",
        create_update_sql_container_parameters={
            "location": "West US",
            "properties": {
                "options": {},
                "resource": {
                    "clientEncryptionPolicy": {
                        "includedPaths": [
                            {
                                "clientEncryptionKeyId": "keyId",
                                "encryptionAlgorithm": "AEAD_AES_256_CBC_HMAC_SHA256",
                                "encryptionType": "Deterministic",
                                "path": "/path",
                            }
                        ],
                        "policyFormatVersion": 2,
                    },
                    "computedProperties": [{"name": "cp_lowerName", "query": "SELECT VALUE LOWER(c.name) FROM c"}],
                    "conflictResolutionPolicy": {"conflictResolutionPath": "/path", "mode": "LastWriterWins"},
                    "defaultTtl": 100,
                    "fullTextPolicy": {
                        "defaultLanguage": "1033",
                        "fullTextPaths": [
                            {"language": "en-US", "path": "/ftPath1"},
                            {"language": "fr-FR", "path": "/ftPath2"},
                            {"language": "de-DE", "path": "/ftPath3"},
                        ],
                    },
                    "id": "containerName",
                    "indexingPolicy": {
                        "automatic": True,
                        "excludedPaths": [],
                        "includedPaths": [
                            {
                                "indexes": [
                                    {"dataType": "String", "kind": "Range", "precision": -1},
                                    {"dataType": "Number", "kind": "Range", "precision": -1},
                                ],
                                "path": "/*",
                            }
                        ],
                        "indexingMode": "consistent",
                        "vectorIndexes": [
                            {"path": "/vectorPath1", "type": "flat"},
                            {"path": "/vectorPath2", "type": "quantizedFlat"},
                            {"path": "/vectorPath3", "type": "diskANN"},
                        ],
                    },
                    "partitionKey": {"kind": "Hash", "paths": ["/AccountNumber"]},
                    "uniqueKeyPolicy": {"uniqueKeys": [{"paths": ["/testPath"]}]},
                    "vectorEmbeddingPolicy": {
                        "vectorEmbeddings": [
                            {
                                "dataType": "float32",
                                "dimensions": 400,
                                "distanceFunction": "euclidean",
                                "path": "/vectorPath1",
                            },
                            {
                                "dataType": "uint8",
                                "dimensions": 512,
                                "distanceFunction": "cosine",
                                "path": "/vectorPath2",
                            },
                            {
                                "dataType": "int8",
                                "dimensions": 512,
                                "distanceFunction": "dotproduct",
                                "path": "/vectorPath3",
                            },
                        ]
                    },
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBSqlContainerCreateUpdate.json
if __name__ == "__main__":
    main()
