from azure.identity import DefaultAzureCredential

from azure.mgmt.relay import RelayAPIMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-relay
# USAGE
    python relay_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RelayAPIMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.wcf_relays.delete(
        resource_group_name="resourcegroup",
        namespace_name="example-RelayNamespace-01",
        relay_name="example-Relay-wcf-01",
    )


# x-ms-original-file: 2024-01-01/Relay/RelayDelete.json
if __name__ == "__main__":
    main()
