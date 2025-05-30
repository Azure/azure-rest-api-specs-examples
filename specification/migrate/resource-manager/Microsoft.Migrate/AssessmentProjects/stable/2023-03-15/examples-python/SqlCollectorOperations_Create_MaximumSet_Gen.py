from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python sql_collector_operations_create_maximum_set_gen.py

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

    response = client.sql_collector_operations.begin_create(
        resource_group_name="rgmigrate",
        project_name="fci-test6904project",
        collector_name="fci-test0c1esqlsitecollector",
        resource={
            "properties": {
                "agentProperties": {
                    "id": "630da710-4d44-41f7-a189-72fe3db5502b-agent",
                    "lastHeartbeatUtc": None,
                    "spnDetails": {
                        "applicationId": "db9c4c3d-477c-4d5a-817b-318276713565",
                        "audience": "db9c4c3d-477c-4d5a-817b-318276713565",
                        "authority": "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47",
                        "objectId": "e50236ad-ad07-47d4-af71-ed7b52d200d5",
                        "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
                    },
                    "version": None,
                },
                "discoverySiteId": "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.OffAzure/MasterSites/fci-ankit-test6065mastersite/SqlSites/fci-ankit-test6065sqlsites",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/SqlCollectorOperations_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
