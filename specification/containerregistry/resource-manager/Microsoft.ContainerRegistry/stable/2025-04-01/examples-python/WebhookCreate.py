from azure.identity import DefaultAzureCredential

from azure.mgmt.containerregistry import ContainerRegistryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerregistry
# USAGE
    python webhook_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerRegistryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.webhooks.begin_create(
        resource_group_name="myResourceGroup",
        registry_name="myRegistry",
        webhook_name="myWebhook",
        webhook_create_parameters={
            "location": "westus",
            "properties": {
                "actions": ["push"],
                "customHeaders": {"Authorization": "******"},
                "scope": "myRepository",
                "serviceUri": "http://myservice.com",
                "status": "enabled",
            },
            "tags": {"key": "value"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/WebhookCreate.json
if __name__ == "__main__":
    main()
