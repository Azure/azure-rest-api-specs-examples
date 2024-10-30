from azure.identity import DefaultAzureCredential

from azure.mgmt.resourcehealth import ResourceHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcehealth
# USAGE
    python operations_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    print(response)


# x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Operations_List.json
if __name__ == "__main__":
    main()
