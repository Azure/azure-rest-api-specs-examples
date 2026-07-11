from azure.identity import DefaultAzureCredential

from azure.mgmt.botservice import AzureBotServiceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-botservice
# USAGE
    python get_host_settings.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureBotServiceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.host_settings.get()
    print(response)


# x-ms-original-file: 2023-09-15-preview/GetHostSettings.json
if __name__ == "__main__":
    main()
