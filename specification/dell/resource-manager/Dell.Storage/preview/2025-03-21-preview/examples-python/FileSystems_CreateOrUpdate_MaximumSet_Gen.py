from azure.identity import DefaultAzureCredential

from azure.mgmt.dellstorage import DellStorageMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dellstorage
# USAGE
    python file_systems_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DellStorageMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.file_systems.begin_create_or_update(
        resource_group_name="rgDell",
        filesystem_name="abcd",
        resource={
            "identity": {"type": "UserAssigned", "userAssignedIdentities": {"key7644": {}}},
            "location": "cvbmsqftppe",
            "properties": {
                "delegatedSubnetCidr": "10.0.0.1/24",
                "delegatedSubnetId": "rqkpvczbtqcxiaivtbuixblb",
                "dellReferenceNumber": "fhewkj",
                "encryption": {
                    "encryptionIdentityProperties": {
                        "identityResourceId": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}",
                        "identityType": "UserAssigned",
                    },
                    "encryptionType": "Customer-managed keys (CMK)",
                    "keyUrl": "https://contoso.com/keyurl/keyVersion",
                },
                "marketplace": {
                    "marketplaceSubscriptionId": "mvjcxwndudbylynme",
                    "marketplaceSubscriptionStatus": "PendingFulfillmentStart",
                    "offerId": "bcganbkmvznyqfnvhjuag",
                    "planId": "eekvwfndjoxijeasksnt",
                    "planName": "planeName",
                    "privateOfferId": "privateOfferId",
                    "publisherId": "trdzykoeskmcwpo",
                },
                "oneFsUrl": "oneFsUrl",
                "smartConnectFqdn": "fqdn",
                "user": {"email": "jwogfgznmjabdbcjcljjlkxdpc"},
            },
            "tags": {"key7594": "sfkwapubiurgedzveido"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-03-21-preview/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
