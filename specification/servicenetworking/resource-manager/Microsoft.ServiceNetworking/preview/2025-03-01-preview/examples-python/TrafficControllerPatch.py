from azure.identity import DefaultAzureCredential

from azure.mgmt.servicenetworking import ServiceNetworkingMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicenetworking
# USAGE
    python traffic_controller_patch.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceNetworkingMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.traffic_controller_interface.update(
        resource_group_name="rg1",
        traffic_controller_name="tc1",
        properties={"tags": {"key1": "value1"}},
    )
    print(response)


# x-ms-original-file: 2025-03-01-preview/TrafficControllerPatch.json
if __name__ == "__main__":
    main()
