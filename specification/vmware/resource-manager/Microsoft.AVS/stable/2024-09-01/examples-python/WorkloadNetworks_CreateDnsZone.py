from azure.identity import DefaultAzureCredential

from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python workload_networks_create_dns_zone.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AVSClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.workload_networks.begin_create_dns_zone(
        resource_group_name="group1",
        private_cloud_name="cloud1",
        dns_zone_id="dnsZone1",
        workload_network_dns_zone={
            "properties": {
                "displayName": "dnsZone1",
                "dnsServerIps": ["1.1.1.1"],
                "domain": [],
                "revision": 1,
                "sourceIp": "8.8.8.8",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2024-09-01/examples/WorkloadNetworks_CreateDnsZone.json
if __name__ == "__main__":
    main()
