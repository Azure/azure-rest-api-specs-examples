from azure.identity import DefaultAzureCredential
from azure.mgmt.networkfunction import TrafficCollectorMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkfunction
# USAGE
    python collector_policy_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TrafficCollectorMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.collector_policies.get(
        resource_group_name="rg1",
        azure_traffic_collector_name="atc",
        collector_policy_name="cp1",
    )
    print(response)


# x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyGet.json
if __name__ == "__main__":
    main()
