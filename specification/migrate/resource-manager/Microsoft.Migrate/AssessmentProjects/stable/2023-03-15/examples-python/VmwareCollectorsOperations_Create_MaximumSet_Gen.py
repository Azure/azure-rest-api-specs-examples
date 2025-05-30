from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python vmware_collectors_operations_create_maximum_set_gen.py

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

    response = client.vmware_collectors_operations.begin_create(
        resource_group_name="ayagrawRG",
        project_name="app18700project",
        vm_ware_collector_name="Vmware2258collector",
        resource={
            "properties": {
                "agentProperties": {
                    "id": "fe243486-3318-41fa-aaba-c48b5df75308",
                    "lastHeartbeatUtc": "2022-03-29T12:10:08.9167289Z",
                    "spnDetails": {
                        "applicationId": "82b3e452-c0e8-4662-8347-58282925ae84",
                        "audience": "82b3e452-c0e8-4662-8347-58282925ae84",
                        "authority": "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47",
                        "objectId": "3fc89111-1405-4938-9214-37aa4739401d",
                        "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
                    },
                    "version": "1.0.8.383",
                },
                "discoverySiteId": "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/VMwareSites/Vmware2744site",
                "provisioningState": "Succeeded",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/VmwareCollectorsOperations_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
