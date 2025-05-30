from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridconnectivity import HybridConnectivityMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridconnectivity
# USAGE
    python service_configurations_put_wac.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridConnectivityMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.service_configurations.create_orupdate(
        resource_uri="subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default",
        endpoint_name="default",
        service_configuration_name="WAC",
        service_configuration_resource={"properties": {"port": 6516, "serviceName": "WAC"}},
    )
    print(response)


# x-ms-original-file: 2024-12-01/ServiceConfigurationsPutWAC.json
if __name__ == "__main__":
    main()
