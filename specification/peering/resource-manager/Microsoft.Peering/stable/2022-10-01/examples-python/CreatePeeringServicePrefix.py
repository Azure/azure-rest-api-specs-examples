from azure.identity import DefaultAzureCredential
from azure.mgmt.peering import PeeringManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-peering
# USAGE
    python create_or_update_a_prefix_for_the_peering_service.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PeeringManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId",
    )

    response = client.prefixes.create_or_update(
        resource_group_name="rgName",
        peering_service_name="peeringServiceName",
        prefix_name="peeringServicePrefixName",
        peering_service_prefix={
            "properties": {
                "peeringServicePrefixKey": "00000000-0000-0000-0000-000000000000",
                "prefix": "192.168.1.0/24",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreatePeeringServicePrefix.json
if __name__ == "__main__":
    main()
