from azure.identity import DefaultAzureCredential
from azure.mgmt.portal import Portal

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-portal
# USAGE
    python get_list_of_tenant_configurations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = Portal(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.tenant_configurations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/TenantConfiguration/GetListOfTenantConfigurations.json
if __name__ == "__main__":
    main()
