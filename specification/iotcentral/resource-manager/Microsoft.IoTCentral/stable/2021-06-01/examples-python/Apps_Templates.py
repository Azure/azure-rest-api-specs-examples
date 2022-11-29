from azure.identity import DefaultAzureCredential
from azure.mgmt.iotcentral import IotCentralClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotcentral
# USAGE
    python apps_templates.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IotCentralClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.apps.list_templates()
    for item in response:
        print(item)


# x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Templates.json
if __name__ == "__main__":
    main()
