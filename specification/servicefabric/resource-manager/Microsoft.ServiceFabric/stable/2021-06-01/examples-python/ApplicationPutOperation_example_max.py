from azure.identity import DefaultAzureCredential
from azure.mgmt.servicefabric import ServiceFabricManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicefabric
# USAGE
    python application_put_operation_example_max.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceFabricManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.applications.begin_create_or_update(
        resource_group_name="resRg",
        cluster_name="myCluster",
        application_name="myApp",
        parameters={
            "properties": {
                "maximumNodes": 3,
                "metrics": [
                    {"maximumCapacity": 3, "name": "metric1", "reservationCapacity": 1, "totalApplicationCapacity": 5}
                ],
                "minimumNodes": 1,
                "parameters": {"param1": "value1"},
                "removeApplicationCapacity": False,
                "typeName": "myAppType",
                "typeVersion": "1.0",
                "upgradePolicy": {
                    "applicationHealthPolicy": {
                        "considerWarningAsError": True,
                        "defaultServiceTypeHealthPolicy": {
                            "maxPercentUnhealthyPartitionsPerService": 0,
                            "maxPercentUnhealthyReplicasPerPartition": 0,
                            "maxPercentUnhealthyServices": 0,
                        },
                        "maxPercentUnhealthyDeployedApplications": 0,
                    },
                    "forceRestart": False,
                    "rollingUpgradeMonitoringPolicy": {
                        "failureAction": "Rollback",
                        "healthCheckRetryTimeout": "00:10:00",
                        "healthCheckStableDuration": "00:05:00",
                        "healthCheckWaitDuration": "00:02:00",
                        "upgradeDomainTimeout": "1.06:00:00",
                        "upgradeTimeout": "01:00:00",
                    },
                    "upgradeMode": "Monitored",
                    "upgradeReplicaSetCheckTimeout": "01:00:00",
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_max.json
if __name__ == "__main__":
    main()
