from azure.identity import DefaultAzureCredential

from azure.mgmt.batch import BatchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-batch
# USAGE
    python pool_create_virtual_machine_configuration_service_artifact_reference.py

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
                        "serviceArtifactReference": {
                            "id": "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/vmArtifactsProfile"
                        },
                        "windowsConfiguration": {"enableAutomaticUpdates": False},
                    }
                },
                "scaleSettings": {"fixedScale": {"targetDedicatedNodes": 2, "targetLowPriorityNodes": 0}},
                "upgradePolicy": {"automaticOSUpgradePolicy": {"enableAutomaticOSUpgrade": True}, "mode": "automatic"},
                "vmSize": "Standard_d4s_v3",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_VirtualMachineConfiguration_ServiceArtifactReference.json
if __name__ == "__main__":
    main()
