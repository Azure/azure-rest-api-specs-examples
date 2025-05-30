from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python assessed_machines_operations_get_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MigrationAssessmentMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="D8E1C413-E65F-40C0-8A7E-743D6B7A6AE9",
    )

    response = client.assessed_machines_operations.get(
        resource_group_name="rgopenapi",
        project_name="pavqtntysjn",
        group_name="smawqdmhfngray",
        assessment_name="qjlumxyqsitd",
        assessed_machine_name="oqxjeheiipjmuo",
    )
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessedMachinesOperations_Get_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
