from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python hyperv_collectors_operations_create_maximum_set_gen.py

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

    response = client.hyperv_collectors_operations.begin_create(
        resource_group_name="ayagrawRG",
        project_name="app18700project",
        hyperv_collector_name="test-697cecollector",
        resource={
            "properties": {
                "agentProperties": {
                    "id": "12f1d90f-b3fa-4926-8893-e56803a09af0",
                    "lastHeartbeatUtc": "2022-07-07T14:25:35.708325Z",
                    "spnDetails": {
                        "applicationId": "e3bd6eaa-980b-40ae-a30e-2a5069ba097c",
                        "audience": "e3bd6eaa-980b-40ae-a30e-2a5069ba097c",
                        "authority": "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47",
                        "objectId": "01b9f9e2-2d82-414c-adaa-09ce259b6b44",
                        "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
                    },
                    "version": "2.0.1993.19",
                },
                "discoverySiteId": "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/HyperVSites/test-60527site",
                "provisioningState": "Succeeded",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/HypervCollectorsOperations_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
