from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningcompute import MachineLearningComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningcompute
# USAGE
    python machine_learning_compute_list_available_operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.machine_learning_compute.list_available_operations()
    print(response)


# x-ms-original-file: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/MachineLearningCompute_ListAvailableOperations.json
if __name__ == "__main__":
    main()
