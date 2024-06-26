from azure.identity import DefaultAzureCredential
from azure.mgmt.networkanalytics import NetworkAnalyticsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkanalytics
# USAGE
    python data_products_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkAnalyticsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-00000000000",
    )

    response = client.data_products.begin_create(
        resource_group_name="aoiresourceGroupName",
        data_product_name="dataproduct01",
        resource={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {}
                },
            },
            "location": "eastus",
            "properties": {
                "consumptionEndpoints": {},
                "currentMinorVersion": "1.0.1",
                "customerEncryptionKey": {
                    "keyName": "keyName",
                    "keyVaultUri": "https://KeyVault.vault.azure.net",
                    "keyVersion": "keyVersion",
                },
                "customerManagedKeyEncryptionEnabled": "Enabled",
                "majorVersion": "1.0.0",
                "managedResourceGroupConfiguration": {"location": "eastus", "name": "managedResourceGroupName"},
                "networkacls": {
                    "allowedQueryIpRangeList": ["1.1.1.1"],
                    "defaultAction": "Allow",
                    "ipRules": [{"action": "Allow", "value": "1.1.1.1"}],
                    "virtualNetworkRule": [
                        {
                            "action": "Allow",
                            "id": "/subscriptions/subscriptionId/resourcegroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/virtualNetworkName/subnets/subnetName",
                            "state": "",
                        }
                    ],
                },
                "owners": ["abc@micros.com"],
                "privateLinksEnabled": "Disabled",
                "product": "MCC",
                "provisioningState": "Succeeded",
                "publicNetworkAccess": "Enabled",
                "publisher": "Microsoft",
                "purviewAccount": "testpurview",
                "purviewCollection": "134567890",
                "redundancy": "Disabled",
            },
            "tags": {"userSpecifiedKeyName": "userSpecifiedKeyValue"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
