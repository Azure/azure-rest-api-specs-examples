from azure.identity import DefaultAzureCredential
from azure.mgmt.defendereasm import EasmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-defendereasm
# USAGE
    python labels_delete.py

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

    response = client.labels.begin_delete(
        resource_group_name="dummyrg",
        workspace_name="ThisisaWorkspace",
        label_name="ThisisaLabel",
    ).result()
    print(response)


# x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2022-04-01-preview/examples/Labels_Delete.json
if __name__ == "__main__":
    main()
