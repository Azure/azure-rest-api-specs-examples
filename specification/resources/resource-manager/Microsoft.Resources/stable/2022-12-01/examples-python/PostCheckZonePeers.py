from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.resource import SubscriptionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python post_check_zone_peers.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SubscriptionClient(
        credential=DefaultAzureCredential(),
    )

    response = client.subscriptions.check_zone_peers(
        subscription_id="8d65815f-a5b6-402f-9298-045155da7d74",
        parameters={"location": "eastus", "subscriptionIds": ["subscriptions/11111111-1111-1111-1111-111111111111"]},
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/PostCheckZonePeers.json
if __name__ == "__main__":
    main()
