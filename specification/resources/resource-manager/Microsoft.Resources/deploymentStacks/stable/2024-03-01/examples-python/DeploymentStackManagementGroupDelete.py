from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.deploymentstacks import DeploymentStacksClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-deploymentstacks
# USAGE
    python deployment_stack_management_group_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DeploymentStacksClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.deployment_stacks.begin_delete_at_management_group(
        management_group_id="myMg",
        deployment_stack_name="simpleDeploymentStack",
    ).result()


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/deploymentStacks/stable/2024-03-01/examples/DeploymentStackManagementGroupDelete.json
if __name__ == "__main__":
    main()
