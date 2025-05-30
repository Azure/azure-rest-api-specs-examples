from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python groups_operations_update_machines_maximum_set_gen.py

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

    response = client.groups_operations.begin_update_machines(
        resource_group_name="ayagrawrg",
        project_name="app18700project",
        group_name="kuchatur-test",
        body={
            "eTag": "*",
            "properties": {
                "machines": [
                    "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/machines/18895660-c5e5-4247-8cfc-cd24e1fe57f3"
                ],
                "operationType": "Add",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/GroupsOperations_UpdateMachines_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
