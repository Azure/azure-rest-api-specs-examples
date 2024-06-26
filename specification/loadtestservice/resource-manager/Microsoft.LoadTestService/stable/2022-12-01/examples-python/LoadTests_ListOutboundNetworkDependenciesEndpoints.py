from azure.identity import DefaultAzureCredential
from azure.mgmt.loadtesting import LoadTestMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loadtesting
# USAGE
    python load_tests_list_outbound_network_dependencies_endpoints.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LoadTestMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.load_tests.list_outbound_network_dependencies_endpoints(
        resource_group_name="default-azureloadtest-japaneast",
        load_test_name="sampleloadtest",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_ListOutboundNetworkDependenciesEndpoints.json
if __name__ == "__main__":
    main()
