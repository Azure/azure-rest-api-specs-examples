from azure.identity import DefaultAzureCredential

from azure.mgmt.azurestackhci import AzureStackHCIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestackhci
# USAGE
    python put_extension.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackHCIClient(
        credential=DefaultAzureCredential(),
        subscription_id="fd3c3665-1729-4b7b-9a38-238e83b0f98b",
    )

    response = client.extensions.begin_create(
        resource_group_name="test-rg",
        cluster_name="myCluster",
        arc_setting_name="default",
        extension_name="MicrosoftMonitoringAgent",
        extension={
            "properties": {
                "extensionParameters": {
                    "enableAutomaticUpgrade": False,
                    "protectedSettings": {"workspaceKey": "xx"},
                    "publisher": "Microsoft.Compute",
                    "settings": {"workspaceId": "xx"},
                    "type": "MicrosoftMonitoringAgent",
                    "typeHandlerVersion": "1.10",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutExtension.json
if __name__ == "__main__":
    main()
