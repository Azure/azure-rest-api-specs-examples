from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python update_workspace.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.workspaces.begin_update(
        resource_group_name="resourceGroup1",
        workspace_name="workspace1",
        workspace_patch_info={
            "identity": {"type": "SystemAssigned"},
            "properties": {
                "encryption": {"cmk": {"key": {"keyVaultUrl": "https://vault.azure.net/keys/key1", "name": "default"}}},
                "managedVirtualNetworkSettings": {
                    "allowedAadTenantIdsForLinking": ["740239CE-A25B-485B-86A0-262F29F6EBDB"],
                    "linkedAccessCheckOnTargetResource": False,
                    "preventDataExfiltration": False,
                },
                "publicNetworkAccess": "Enabled",
                "purviewConfiguration": {
                    "purviewResourceId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"
                },
                "sqlAdministratorLoginPassword": "password",
                "workspaceRepositoryConfiguration": {
                    "accountName": "adifferentacount",
                    "collaborationBranch": "master",
                    "hostName": "",
                    "projectName": "myproject",
                    "repositoryName": "myrepository",
                    "rootFolder": "/",
                    "type": "FactoryGitHubConfiguration",
                },
            },
            "tags": {"key": "value"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
if __name__ == "__main__":
    main()
