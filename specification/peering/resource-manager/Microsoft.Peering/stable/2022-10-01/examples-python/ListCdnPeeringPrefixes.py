from azure.identity import DefaultAzureCredential
from azure.mgmt.peering import PeeringManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-peering
# USAGE
    python list_all_the_cdn_peering_prefixes_advertised_at_a_particular_peering_location.py

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

    response = client.cdn_peering_prefixes.list(
        peering_location="peeringLocation0",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/ListCdnPeeringPrefixes.json
if __name__ == "__main__":
    main()
