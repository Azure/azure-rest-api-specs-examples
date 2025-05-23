from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python server_collectors_operations_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MigrationAssessmentMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="4bd2aa0f-2bd2-4d67-91a8-5a4533d58600",
    )

    response = client.server_collectors_operations.begin_create(
        resource_group_name="ayagrawRG",
        project_name="app18700project",
        server_collector_name="walter389fcollector",
        resource={
            "properties": {
                "agentProperties": {
                    "id": "498e4965-bbb1-47c2-8613-345baff9c509",
                    "lastHeartbeatUtc": None,
                    "spnDetails": {
                        "applicationId": "65153d2f-9afb-44e8-b3ca-1369150b7354",
                        "audience": "65153d2f-9afb-44e8-b3ca-1369150b7354",
                        "authority": "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47",
                        "objectId": "ddde6f96-87c8-420b-9d4d-f16a5090519e",
                        "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
                    },
                    "version": None,
                },
                "discoverySiteId": "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/ServerSites/walter7155site",
                "provisioningState": "Succeeded",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/ServerCollectorsOperations_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
