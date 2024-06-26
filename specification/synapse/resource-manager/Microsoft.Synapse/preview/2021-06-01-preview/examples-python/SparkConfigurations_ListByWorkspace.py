from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python spark_configurations_list_by_workspace.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-12345678abc",
    )

    response = client.spark_configurations.list_by_workspace(
        resource_group_name="exampleResourceGroup",
        workspace_name="exampleWorkspace",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/SparkConfigurations_ListByWorkspace.json
if __name__ == "__main__":
    main()
