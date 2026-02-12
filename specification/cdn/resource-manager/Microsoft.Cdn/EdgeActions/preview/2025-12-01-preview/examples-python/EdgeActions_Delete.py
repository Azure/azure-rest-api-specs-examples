from azure.identity import DefaultAzureCredential

from azure.mgmt.edgeactions import EdgeActionsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-edgeactions
# USAGE
    python edge_actions_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EdgeActionsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.edge_actions.begin_delete(
        resource_group_name="testrg",
        edge_action_name="edgeAction1",
    ).result()


# x-ms-original-file: 2025-12-01-preview/EdgeActions_Delete.json
if __name__ == "__main__":
    main()
