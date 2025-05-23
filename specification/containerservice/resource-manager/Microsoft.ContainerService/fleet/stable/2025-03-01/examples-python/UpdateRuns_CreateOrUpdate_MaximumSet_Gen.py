from azure.identity import DefaultAzureCredential

from azure.mgmt.containerservicefleet import ContainerServiceFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservicefleet
# USAGE
    python update_runs_create_or_update_maximum_set_gen.py

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

    response = client.update_runs.begin_create_or_update(
        resource_group_name="rgfleets",
        fleet_name="fleet1",
        update_run_name="fleet1",
        resource={
            "properties": {
                "autoUpgradeProfileId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/autoUpgradeProfiles/aup1",
                "managedClusterUpdate": {
                    "nodeImageSelection": {"customNodeImageVersions": [{}], "type": "Latest"},
                    "upgrade": {"kubernetesVersion": "1.26.1", "type": "Full"},
                },
                "status": {"nodeImageSelection": {}, "status": {"error": {}, "state": "NotStarted"}},
                "strategy": {
                    "stages": [{"afterStageWaitInSeconds": 3600, "groups": [{"name": "group-a"}], "name": "stage1"}]
                },
                "updateStrategyId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2025-03-01/examples/UpdateRuns_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
