from azure.identity import DefaultAzureCredential

from azure.mgmt.cognitiveservices import CognitiveServicesManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cognitiveservices
# USAGE
    python get_operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CognitiveServicesManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/GetOperations.json
if __name__ == "__main__":
    main()
