from azure.identity import DefaultAzureCredential
from azure.mgmt.automanage import AutomanageClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automanage
# USAGE
    python list_best_practices_by_tenant.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomanageClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.best_practices.list_by_tenant()
    for item in response:
        print(item)


# x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listBestPracticesByTenant.json
if __name__ == "__main__":
    main()
