from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningservices import MachineLearningServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningservices
# USAGE
    python list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.registry_code_containers.list(
        resource_group_name="testrg123",
        registry_name="testregistry",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2023-04-01/examples/Registry/CodeContainer/list.json
if __name__ == "__main__":
    main()
