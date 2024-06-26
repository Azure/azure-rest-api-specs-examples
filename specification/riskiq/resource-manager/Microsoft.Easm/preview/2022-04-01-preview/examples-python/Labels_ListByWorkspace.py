from azure.identity import DefaultAzureCredential
from azure.mgmt.defendereasm import EasmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-defendereasm
# USAGE
    python labels_list_by_workspace.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EasmMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.labels.list_by_workspace(
        resource_group_name="dummyrg",
        workspace_name="ThisisaWorkspace",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2022-04-01-preview/examples/Labels_ListByWorkspace.json
if __name__ == "__main__":
    main()
