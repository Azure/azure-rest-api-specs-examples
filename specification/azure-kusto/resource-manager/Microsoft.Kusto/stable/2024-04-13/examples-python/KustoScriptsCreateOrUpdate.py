from azure.identity import DefaultAzureCredential

from azure.mgmt.kusto import KustoManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kusto
# USAGE
    python kusto_scripts_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KustoManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-123456789098",
    )

    response = client.scripts.begin_create_or_update(
        resource_group_name="kustorptest",
        cluster_name="kustoCluster",
        database_name="KustoDatabase8",
        script_name="kustoScript",
        parameters={
            "properties": {
                "continueOnErrors": True,
                "forceUpdateTag": "2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe",
                "principalPermissionsAction": "RemovePermissionOnScriptCompletion",
                "scriptLevel": "Database",
                "scriptUrl": "https://mysa.blob.core.windows.net/container/script.txt",
                "scriptUrlSasToken": "?sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=********************************",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoScriptsCreateOrUpdate.json
if __name__ == "__main__":
    main()
