from azure.identity import DefaultAzureCredential

from azure.mgmt.elastic import MicrosoftElastic

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elastic
# USAGE
    python open_ai_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftElastic(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.open_ai.create_or_update(
        resource_group_name="myResourceGroup",
        monitor_name="myMonitor",
        integration_name="default",
    )
    print(response)


# x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2024-06-15-preview/examples/OpenAI_CreateOrUpdate.json
if __name__ == "__main__":
    main()
