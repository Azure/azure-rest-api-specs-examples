from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.standbypool import StandbyPoolMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-standbypool
# USAGE
    python standby_container_group_pools_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StandbyPoolMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="8CC31D61-82D7-4B2B-B9DC-6B924DE7D229",
    )

    response = client.standby_container_group_pools.begin_create_or_update(
        resource_group_name="rgstandbypool",
        standby_container_group_pool_name="pool",
        resource={
            "location": "West US",
            "properties": {
                "containerGroupProperties": {
                    "containerGroupProfile": {
                        "id": "/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile",
                        "revision": 1,
                    },
                    "subnetIds": [
                        {
                            "id": "/subscriptions/8cf6c1b6-c80f-437c-87ad-45fbaff54f73/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"
                        }
                    ],
                },
                "elasticityProfile": {"maxReadyCapacity": 688, "refillPolicy": "always"},
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyContainerGroupPools_CreateOrUpdate.json
if __name__ == "__main__":
    main()
