from azure.identity import DefaultAzureCredential

from azure.mgmt.standbypool import StandbyPoolMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-standbypool
# USAGE
    python standby_container_group_pools_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StandbyPoolMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.standby_container_group_pools.update(
        resource_group_name="rgstandbypool",
        standby_container_group_pool_name="pool",
        properties={
            "properties": {
                "containerGroupProperties": {
                    "containerGroupProfile": {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile",
                        "revision": 2,
                    },
                    "subnetIds": [
                        {
                            "id": "/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"
                        }
                    ],
                },
                "elasticityProfile": {"maxReadyCapacity": 1743, "refillPolicy": "always"},
                "zones": ["1", "2", "3"],
            },
            "tags": {},
        },
    )
    print(response)


# x-ms-original-file: 2025-03-01/StandbyContainerGroupPools_Update.json
if __name__ == "__main__":
    main()
