from azure.identity import DefaultAzureCredential
from azure.mgmt.confluent import ConfluentManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confluent
# USAGE
    python organization_operations_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConfluentManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.organization_operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/OrganizationOperations_List.json
if __name__ == "__main__":
    main()
