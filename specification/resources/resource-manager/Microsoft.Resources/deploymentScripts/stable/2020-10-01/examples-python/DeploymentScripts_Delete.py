from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.deploymentscripts import DeploymentScriptsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-deploymentscripts
# USAGE
    python deployment_scripts_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DeploymentScriptsClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.deployment_scripts.delete(
        resource_group_name="script-rg",
        script_name="MyDeploymentScript",
    )


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/deploymentScripts/stable/2020-10-01/examples/DeploymentScripts_Delete.json
if __name__ == "__main__":
    main()
