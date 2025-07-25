from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.deploymentscripts import DeploymentScriptsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-deploymentscripts
# USAGE
    python deployment_scripts_create.py

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

    response = client.deployment_scripts.begin_create(
        resource_group_name="script-rg",
        script_name="MyDeploymentScript",
        deployment_script={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai": {}
                },
            },
            "kind": "AzurePowerShell",
            "location": "westus",
            "properties": {
                "arguments": "-Location 'westus' -Name \"*rg2\"",
                "azPowerShellVersion": "1.7.0",
                "cleanupPreference": "Always",
                "retentionInterval": "PT7D",
                "scriptContent": "Param([string]$Location,[string]$Name) $deploymentScriptOutputs['test'] = 'value' Get-AzResourceGroup -Location $Location -Name $Name",
                "supportingScriptUris": ["https://uri1.to.supporting.script", "https://uri2.to.supporting.script"],
                "timeout": "PT1H",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/deploymentScripts/stable/2020-10-01/examples/DeploymentScripts_Create.json
if __name__ == "__main__":
    main()
