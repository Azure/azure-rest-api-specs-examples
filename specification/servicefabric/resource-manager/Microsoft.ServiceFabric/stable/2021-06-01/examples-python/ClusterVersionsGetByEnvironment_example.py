from azure.identity import DefaultAzureCredential
from azure.mgmt.servicefabric import ServiceFabricManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicefabric
# USAGE
    python cluster_versions_get_by_environment_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceFabricManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.cluster_versions.get_by_environment(
        location="eastus",
        environment="Windows",
        cluster_version="6.1.480.9494",
    )
    print(response)


# x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGetByEnvironment_example.json
if __name__ == "__main__":
    main()
