from azure.identity import DefaultAzureCredential

from azure.mgmt.onlineexperimentation import OnlineExperimentationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-onlineexperimentation
# USAGE
    python online_experimentation_workspaces_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OnlineExperimentationMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.online_experimentation_workspaces.begin_update(
        resource_group_name="res9871",
        workspace_name="expworkspace3",
        properties={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
                },
            },
            "tags": {"newKey": "newVal"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_Update.json
if __name__ == "__main__":
    main()
