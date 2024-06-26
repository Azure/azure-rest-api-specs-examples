from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python enable_interactivequery_integration_runtimes.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="2d03866b-587b-4e1f-a2fe-0a55958c655e",
    )

    response = client.get.integration_runtime_enable_interactivequery(
        resource_group_name="drage-felles-prod-rg",
        workspace_name="felles-prod-synapse-workspace",
        integration_runtime_name="SSIS-intergrationRuntime-Drage",
        integration_runtime_operation_id="5752dcdf918e4aecb941245ddf6ebb83",
    )
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/EnableInteractivequery_IntegrationRuntimes.json
if __name__ == "__main__":
    main()
