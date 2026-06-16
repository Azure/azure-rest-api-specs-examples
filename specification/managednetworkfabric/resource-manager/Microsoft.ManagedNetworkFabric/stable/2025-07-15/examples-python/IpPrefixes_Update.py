from azure.identity import DefaultAzureCredential

from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python ip_prefixes_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedNetworkFabricMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.ip_prefixes.begin_update(
        resource_group_name="example-rg",
        ip_prefix_name="example-ipPrefix",
        body={
            "properties": {
                "annotation": "annotation",
                "ipPrefixRules": [
                    {
                        "action": "Permit",
                        "condition": "GreaterThanOrEqualTo",
                        "networkPrefix": "10.10.10.10/30",
                        "sequenceNumber": 4155123341,
                        "subnetMaskLength": "10",
                    }
                ],
            },
            "tags": {"KeyId": "KeyValue"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-07-15/IpPrefixes_Update.json
if __name__ == "__main__":
    main()
