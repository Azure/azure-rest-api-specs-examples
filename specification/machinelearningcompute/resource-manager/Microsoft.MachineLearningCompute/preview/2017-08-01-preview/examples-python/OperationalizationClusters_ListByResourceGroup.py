from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningcompute import MachineLearningComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningcompute
# USAGE
    python list_operationalization_clusters_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.operationalization_clusters.list_by_resource_group(
        resource_group_name="myResourceGroup",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/OperationalizationClusters_ListByResourceGroup.json
if __name__ == "__main__":
    main()
