from azure.identity import DefaultAzureCredential

from azure.mgmt.privatedns import PrivateDnsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-privatedns
# USAGE
    python record_set_srv_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PrivateDnsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.record_sets.list_by_type(
        resource_group_name="resourceGroup1",
        private_zone_name="privatezone1.com",
        record_type="SRV",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/RecordSetSRVList.json
if __name__ == "__main__":
    main()
