from azure.identity import DefaultAzureCredential

from azure.mgmt.containerservicefleet import ContainerServiceFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservicefleet
# USAGE
    python update_strategies_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerServiceFleetMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.fleet_update_strategies.begin_create_or_update(
        resource_group_name="rg1",
        fleet_name="fleet1",
        update_strategy_name="strartegy1",
        resource={
            "properties": {
                "strategy": {
                    "stages": [{"afterStageWaitInSeconds": 3600, "groups": [{"name": "group-a"}], "name": "stage1"}]
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/UpdateStrategies_CreateOrUpdate.json
if __name__ == "__main__":
    main()
