from azure.identity import DefaultAzureCredential
from azure.mgmt.labservices import ManagedLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-labservices
# USAGE
    python list_operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedLabsClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabServices/listOperations.json
if __name__ == "__main__":
    main()
