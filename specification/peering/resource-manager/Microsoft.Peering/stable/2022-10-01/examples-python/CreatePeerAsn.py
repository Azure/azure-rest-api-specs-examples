from azure.identity import DefaultAzureCredential
from azure.mgmt.peering import PeeringManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-peering
# USAGE
    python create_a_peer_asn.py

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

    response = client.peer_asns.create_or_update(
        peer_asn_name="peerAsnName",
        peer_asn={
            "properties": {
                "peerAsn": 65000,
                "peerContactDetail": [
                    {"email": "noc@contoso.com", "phone": "+1 (234) 567-8999", "role": "Noc"},
                    {"email": "abc@contoso.com", "phone": "+1 (234) 567-8900", "role": "Policy"},
                    {"email": "xyz@contoso.com", "phone": "+1 (234) 567-8900", "role": "Technical"},
                ],
                "peerName": "Contoso",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreatePeerAsn.json
if __name__ == "__main__":
    main()
