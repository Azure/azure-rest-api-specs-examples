from azure.identity import DefaultAzureCredential
from azure.mgmt.customproviders import Customproviders

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-customproviders
# USAGE
    python list_the_custom_providers_operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = Customproviders(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/operationsList.json
if __name__ == "__main__":
    main()
