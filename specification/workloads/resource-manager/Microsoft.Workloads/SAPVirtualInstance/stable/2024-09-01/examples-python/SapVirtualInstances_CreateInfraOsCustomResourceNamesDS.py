from azure.identity import DefaultAzureCredential

from azure.mgmt.workloadssapvirtualinstance import WorkloadsSapVirtualInstanceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadssapvirtualinstance
# USAGE
    python sap_virtual_instances_create_infra_os_custom_resource_names_ds.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsSapVirtualInstanceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.sap_virtual_instances.begin_create(
        resource_group_name="test-rg",
        sap_virtual_instance_name="X00",
        resource={
            "location": "westcentralus",
            "properties": {
                "configuration": {
                    "appLocation": "eastus",
                    "configurationType": "DeploymentWithOSConfig",
                    "infrastructureConfiguration": {
                        "appResourceGroup": "X00-RG",
                        "applicationServer": {
                            "instanceCount": 6,
                            "subnetId": "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
                            "virtualMachineConfiguration": {
                                "imageReference": {
                                    "offer": "RHEL-SAP",
                                    "publisher": "RedHat",
                                    "sku": "84sapha-gen2",
                                    "version": "latest",
                                },
                                "osProfile": {
                                    "adminUsername": "{your-username}",
                                    "osConfiguration": {
                                        "disablePasswordAuthentication": True,
                                        "osType": "Linux",
                                        "sshKeyPair": {"privateKey": "xyz", "publicKey": "abc"},
                                    },
                                },
                                "vmSize": "Standard_E32ds_v4",
                            },
                        },
                        "centralServer": {
                            "instanceCount": 1,
                            "subnetId": "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
                            "virtualMachineConfiguration": {
                                "imageReference": {
                                    "offer": "RHEL-SAP",
                                    "publisher": "RedHat",
                                    "sku": "84sapha-gen2",
                                    "version": "latest",
                                },
                                "osProfile": {
                                    "adminUsername": "{your-username}",
                                    "osConfiguration": {
                                        "disablePasswordAuthentication": True,
                                        "osType": "Linux",
                                        "sshKeyPair": {"privateKey": "xyz", "publicKey": "abc"},
                                    },
                                },
                                "vmSize": "Standard_E16ds_v4",
                            },
                        },
                        "customResourceNames": {
                            "applicationServer": {
                                "availabilitySetName": "appAvSet",
                                "virtualMachines": [
                                    {
                                        "dataDiskNames": {"default": ["app0disk0"]},
                                        "hostName": "apphostName0",
                                        "networkInterfaces": [{"networkInterfaceName": "appnic0"}],
                                        "osDiskName": "app0osdisk",
                                        "vmName": "appvm0",
                                    },
                                    {
                                        "dataDiskNames": {"default": ["app1disk0"]},
                                        "hostName": "apphostName1",
                                        "networkInterfaces": [{"networkInterfaceName": "appnic1"}],
                                        "osDiskName": "app1osdisk",
                                        "vmName": "appvm1",
                                    },
                                ],
                            },
                            "centralServer": {
                                "virtualMachines": [
                                    {
                                        "dataDiskNames": {"default": ["ascsdisk0"]},
                                        "hostName": "ascshostName",
                                        "networkInterfaces": [{"networkInterfaceName": "ascsnic"}],
                                        "osDiskName": "ascsosdisk",
                                        "vmName": "ascsvm",
                                    }
                                ]
                            },
                            "databaseServer": {
                                "virtualMachines": [
                                    {
                                        "dataDiskNames": {
                                            "hanaData": ["hanadata0", "hanadata1"],
                                            "hanaLog": ["hanalog0", "hanalog1", "hanalog2"],
                                            "hanaShared": ["hanashared0", "hanashared1"],
                                            "usrSap": ["usrsap0"],
                                        },
                                        "hostName": "dbhostName",
                                        "networkInterfaces": [{"networkInterfaceName": "dbnic"}],
                                        "osDiskName": "dbosdisk",
                                        "vmName": "dbvm",
                                    }
                                ]
                            },
                            "namingPatternType": "FullResourceName",
                            "sharedStorage": {
                                "sharedStorageAccountName": "storageacc",
                                "sharedStorageAccountPrivateEndPointName": "peForxNFS",
                            },
                        },
                        "databaseServer": {
                            "databaseType": "HANA",
                            "instanceCount": 1,
                            "subnetId": "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet",
                            "virtualMachineConfiguration": {
                                "imageReference": {
                                    "offer": "RHEL-SAP",
                                    "publisher": "RedHat",
                                    "sku": "84sapha-gen2",
                                    "version": "latest",
                                },
                                "osProfile": {
                                    "adminUsername": "{your-username}",
                                    "osConfiguration": {
                                        "disablePasswordAuthentication": True,
                                        "osType": "Linux",
                                        "sshKeyPair": {"privateKey": "xyz", "publicKey": "abc"},
                                    },
                                },
                                "vmSize": "Standard_M32ts",
                            },
                        },
                        "deploymentType": "ThreeTier",
                    },
                    "osSapConfiguration": {"sapFqdn": "xyz.test.com"},
                },
                "environment": "Prod",
                "sapProduct": "S4HANA",
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraOsCustomResourceNamesDS.json
if __name__ == "__main__":
    main()
