from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python assessment_options_operations_list_by_assessment_project_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MigrationAssessmentMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="A926B99C-7F4C-4556-871E-20CB8C6ADB56",
    )

    response = client.assessment_options_operations.list_by_assessment_project(
        resource_group_name="rgmigrate",
        project_name="fhodvffhuoqwbysrrqbizete",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentOptionsOperations_ListByAssessmentProject_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
