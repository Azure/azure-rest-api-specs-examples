from azure.identity import DefaultAzureCredential

from azure.mgmt.netapp import NetAppManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-netapp
# USAGE
    python volume_groups_create_sap_hana.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetAppManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.volume_groups.begin_create(
        resource_group_name="myRG",
        account_name="account1",
        volume_group_name="group1",
        body={
            "location": "westus",
            "properties": {
                "groupMetaData": {
                    "applicationIdentifier": "SH9",
                    "applicationType": "SAP-HANA",
                    "groupDescription": "Volume group",
                },
                "volumes": [
                    {
                        "name": "test-data-mnt00001",
                        "properties": {
                            "capacityPoolResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
                            "creationToken": "test-data-mnt00001",
                            "exportPolicy": {
                                "rules": [
                                    {
                                        "allowedClients": "0.0.0.0/0",
                                        "cifs": False,
                                        "hasRootAccess": True,
                                        "kerberos5ReadOnly": False,
                                        "kerberos5ReadWrite": False,
                                        "kerberos5iReadOnly": False,
                                        "kerberos5iReadWrite": False,
                                        "kerberos5pReadOnly": False,
                                        "kerberos5pReadWrite": False,
                                        "nfsv3": False,
                                        "nfsv41": True,
                                        "ruleIndex": 1,
                                        "unixReadOnly": True,
                                        "unixReadWrite": True,
                                    }
                                ]
                            },
                            "protocolTypes": ["NFSv4.1"],
                            "proximityPlacementGroup": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
                            "serviceLevel": "Premium",
                            "subnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
                            "throughputMibps": 10,
                            "usageThreshold": 107374182400,
                            "volumeSpecName": "data",
                        },
                    },
                    {
                        "name": "test-log-mnt00001",
                        "properties": {
                            "capacityPoolResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
                            "creationToken": "test-log-mnt00001",
                            "exportPolicy": {
                                "rules": [
                                    {
                                        "allowedClients": "0.0.0.0/0",
                                        "cifs": False,
                                        "hasRootAccess": True,
                                        "kerberos5ReadOnly": False,
                                        "kerberos5ReadWrite": False,
                                        "kerberos5iReadOnly": False,
                                        "kerberos5iReadWrite": False,
                                        "kerberos5pReadOnly": False,
                                        "kerberos5pReadWrite": False,
                                        "nfsv3": False,
                                        "nfsv41": True,
                                        "ruleIndex": 1,
                                        "unixReadOnly": True,
                                        "unixReadWrite": True,
                                    }
                                ]
                            },
                            "protocolTypes": ["NFSv4.1"],
                            "proximityPlacementGroup": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
                            "serviceLevel": "Premium",
                            "subnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
                            "throughputMibps": 10,
                            "usageThreshold": 107374182400,
                            "volumeSpecName": "log",
                        },
                    },
                    {
                        "name": "test-shared",
                        "properties": {
                            "capacityPoolResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
                            "creationToken": "test-shared",
                            "exportPolicy": {
                                "rules": [
                                    {
                                        "allowedClients": "0.0.0.0/0",
                                        "cifs": False,
                                        "hasRootAccess": True,
                                        "kerberos5ReadOnly": False,
                                        "kerberos5ReadWrite": False,
                                        "kerberos5iReadOnly": False,
                                        "kerberos5iReadWrite": False,
                                        "kerberos5pReadOnly": False,
                                        "kerberos5pReadWrite": False,
                                        "nfsv3": False,
                                        "nfsv41": True,
                                        "ruleIndex": 1,
                                        "unixReadOnly": True,
                                        "unixReadWrite": True,
                                    }
                                ]
                            },
                            "protocolTypes": ["NFSv4.1"],
                            "proximityPlacementGroup": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
                            "serviceLevel": "Premium",
                            "subnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
                            "throughputMibps": 10,
                            "usageThreshold": 107374182400,
                            "volumeSpecName": "shared",
                        },
                    },
                    {
                        "name": "test-data-backup",
                        "properties": {
                            "capacityPoolResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
                            "creationToken": "test-data-backup",
                            "exportPolicy": {
                                "rules": [
                                    {
                                        "allowedClients": "0.0.0.0/0",
                                        "cifs": False,
                                        "hasRootAccess": True,
                                        "kerberos5ReadOnly": False,
                                        "kerberos5ReadWrite": False,
                                        "kerberos5iReadOnly": False,
                                        "kerberos5iReadWrite": False,
                                        "kerberos5pReadOnly": False,
                                        "kerberos5pReadWrite": False,
                                        "nfsv3": False,
                                        "nfsv41": True,
                                        "ruleIndex": 1,
                                        "unixReadOnly": True,
                                        "unixReadWrite": True,
                                    }
                                ]
                            },
                            "protocolTypes": ["NFSv4.1"],
                            "proximityPlacementGroup": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
                            "serviceLevel": "Premium",
                            "subnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
                            "throughputMibps": 10,
                            "usageThreshold": 107374182400,
                            "volumeSpecName": "data-backup",
                        },
                    },
                    {
                        "name": "test-log-backup",
                        "properties": {
                            "capacityPoolResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
                            "creationToken": "test-log-backup",
                            "exportPolicy": {
                                "rules": [
                                    {
                                        "allowedClients": "0.0.0.0/0",
                                        "cifs": False,
                                        "hasRootAccess": True,
                                        "kerberos5ReadOnly": False,
                                        "kerberos5ReadWrite": False,
                                        "kerberos5iReadOnly": False,
                                        "kerberos5iReadWrite": False,
                                        "kerberos5pReadOnly": False,
                                        "kerberos5pReadWrite": False,
                                        "nfsv3": False,
                                        "nfsv41": True,
                                        "ruleIndex": 1,
                                        "unixReadOnly": True,
                                        "unixReadWrite": True,
                                    }
                                ]
                            },
                            "protocolTypes": ["NFSv4.1"],
                            "proximityPlacementGroup": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
                            "serviceLevel": "Premium",
                            "subnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
                            "throughputMibps": 10,
                            "usageThreshold": 107374182400,
                            "volumeSpecName": "log-backup",
                        },
                    },
                ],
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/VolumeGroups_Create_SapHana.json
if __name__ == "__main__":
    main()
