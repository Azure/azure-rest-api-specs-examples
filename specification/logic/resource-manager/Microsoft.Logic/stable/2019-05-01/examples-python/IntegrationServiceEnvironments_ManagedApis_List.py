from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python gets_the_integration_service_environment_managed_apis.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="f34b22a3-2202-4fb1-b040-1332bd928c84",
    )

    response = client.integration_service_environment_managed_apis.list(
        resource_group="testResourceGroup",
        integration_service_environment_name="testIntegrationServiceEnvironment",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_List.json
if __name__ == "__main__":
    main()
