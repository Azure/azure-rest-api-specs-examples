from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python delete_an_integration_account_assembly.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.integration_account_assemblies.delete(
        resource_group_name="testResourceGroup",
        integration_account_name="testIntegrationAccount",
        assembly_artifact_name="testAssembly",
    )
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_Delete.json
if __name__ == "__main__":
    main()
