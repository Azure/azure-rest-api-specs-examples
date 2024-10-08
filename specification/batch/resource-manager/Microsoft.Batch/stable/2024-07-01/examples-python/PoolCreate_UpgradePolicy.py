from azure.identity import DefaultAzureCredential

from azure.mgmt.batch import BatchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-batch
# USAGE
    python pool_create_upgrade_policy.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BatchManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.pool.create(
        resource_group_name="default-azurebatch-japaneast",
        account_name="sampleacct",
        pool_name="testpool",
        parameters={
            "properties": {
                "deploymentConfiguration": {
                    "virtualMachineConfiguration": {
                        "imageReference": {
                            "offer": "WindowsServer",
                            "publisher": "MicrosoftWindowsServer",
                            "sku": "2019-datacenter-smalldisk",
                            "version": "latest",
                        },
                        "nodeAgentSkuId": "batch.node.windows amd64",
                        "nodePlacementConfiguration": {"policy": "Zonal"},
                        "windowsConfiguration": {"enableAutomaticUpdates": False},
                    }
                },
                "scaleSettings": {"fixedScale": {"targetDedicatedNodes": 2, "targetLowPriorityNodes": 0}},
                "upgradePolicy": {
                    "automaticOSUpgradePolicy": {
                        "disableAutomaticRollback": True,
                        "enableAutomaticOSUpgrade": True,
                        "osRollingUpgradeDeferral": True,
                        "useRollingUpgradePolicy": True,
                    },
                    "mode": "automatic",
                    "rollingUpgradePolicy": {
                        "enableCrossZoneUpgrade": True,
                        "maxBatchInstancePercent": 20,
                        "maxUnhealthyInstancePercent": 20,
                        "maxUnhealthyUpgradedInstancePercent": 20,
                        "pauseTimeBetweenBatches": "PT0S",
                        "prioritizeUnhealthyInstances": False,
                        "rollbackFailedInstancesOnPolicyBreach": False,
                    },
                },
                "vmSize": "Standard_d4s_v3",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_UpgradePolicy.json
if __name__ == "__main__":
    main()
