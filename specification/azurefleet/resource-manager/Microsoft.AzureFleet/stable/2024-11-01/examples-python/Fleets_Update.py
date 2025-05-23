from azure.identity import DefaultAzureCredential

from azure.mgmt.computefleet import ComputeFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-computefleet
# USAGE
    python fleets_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeFleetMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.fleets.begin_update(
        resource_group_name="rgazurefleet",
        fleet_name="testFleet",
        properties={
            "identity": {"type": "UserAssigned", "userAssignedIdentities": {}},
            "plan": {
                "name": "jwgrcrnrtfoxn",
                "product": "cgopbyvdyqikahwyxfpzwaqk",
                "promotionCode": "naglezezplcaruqogtxnuizslqnnbr",
                "publisher": "iozjbiqqckqm",
                "version": "wa",
            },
            "properties": {
                "computeProfile": {
                    "baseVirtualMachineProfile": {
                        "applicationProfile": {
                            "galleryApplications": [
                                {
                                    "configurationReference": "ulztmiavpojpbpbddgnuuiimxcpau",
                                    "enableAutomaticUpgrade": True,
                                    "order": 5,
                                    "packageReferenceId": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{applicationName}/versions/{versionName}",
                                    "tags": "eyrqjbib",
                                    "treatFailureAsDeploymentFailure": True,
                                }
                            ]
                        },
                        "capacityReservation": {
                            "capacityReservationGroup": {
                                "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
                            }
                        },
                        "diagnosticsProfile": {
                            "bootDiagnostics": {
                                "enabled": True,
                                "storageUri": "http://myStorageAccountName.blob.core.windows.net",
                            }
                        },
                        "extensionProfile": {
                            "extensions": [
                                {
                                    "name": "bndxuxx",
                                    "properties": {
                                        "autoUpgradeMinorVersion": True,
                                        "enableAutomaticUpgrade": True,
                                        "forceUpdateTag": "yhgxw",
                                        "protectedSettings": {},
                                        "protectedSettingsFromKeyVault": {
                                            "secretUrl": "https://myVaultName.vault.azure.net/secrets/secret/mySecretName",
                                            "sourceVault": {
                                                "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
                                            },
                                        },
                                        "provisionAfterExtensions": ["nftzosroolbcwmpupujzqwqe"],
                                        "publisher": "kpxtirxjfprhs",
                                        "settings": {},
                                        "suppressFailures": True,
                                        "type": "pgjilctjjwaa",
                                        "typeHandlerVersion": "zevivcoilxmbwlrihhhibq",
                                    },
                                }
                            ],
                            "extensionsTimeBudget": "mbhjahtdygwgyszdwjtvlvtgchdwil",
                        },
                        "hardwareProfile": {"vmSizeProperties": {"vCPUsAvailable": 16, "vCPUsPerCore": 23}},
                        "licenseType": "v",
                        "networkProfile": {
                            "healthProbe": {
                                "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}"
                            },
                            "networkApiVersion": "2020-11-01",
                            "networkInterfaceConfigurations": [
                                {
                                    "name": "i",
                                    "properties": {
                                        "auxiliaryMode": "None",
                                        "auxiliarySku": "None",
                                        "deleteOption": "Delete",
                                        "disableTcpStateTracking": True,
                                        "dnsSettings": {"dnsServers": ["nxmmfolhclsesu"]},
                                        "enableAcceleratedNetworking": True,
                                        "enableFpga": True,
                                        "enableIPForwarding": True,
                                        "ipConfigurations": [
                                            {
                                                "name": "oezqhkidfhyywlfzwuotilrpbqnjg",
                                                "properties": {
                                                    "applicationGatewayBackendAddressPools": [
                                                        {
                                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/backendAddressPools/{backendAddressPoolName}"
                                                        }
                                                    ],
                                                    "applicationSecurityGroups": [
                                                        {
                                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
                                                        }
                                                    ],
                                                    "loadBalancerBackendAddressPools": [
                                                        {
                                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
                                                        }
                                                    ],
                                                    "loadBalancerInboundNatPools": [
                                                        {
                                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatPools/{inboundNatPoolName}"
                                                        }
                                                    ],
                                                    "primary": True,
                                                    "privateIPAddressVersion": "IPv4",
                                                    "publicIPAddressConfiguration": {
                                                        "name": "fvpqf",
                                                        "properties": {
                                                            "deleteOption": "Delete",
                                                            "dnsSettings": {
                                                                "domainNameLabel": "ukrddzvmorpmfsczjwtbvp",
                                                                "domainNameLabelScope": "TenantReuse",
                                                            },
                                                            "idleTimeoutInMinutes": 9,
                                                            "ipTags": [
                                                                {
                                                                    "ipTagType": "sddgsoemnzgqizale",
                                                                    "tag": "wufmhrjsakbiaetyara",
                                                                }
                                                            ],
                                                            "publicIPAddressVersion": "IPv4",
                                                            "publicIPPrefix": {
                                                                "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}"
                                                            },
                                                        },
                                                        "sku": {"name": "Basic", "tier": "Regional"},
                                                    },
                                                    "subnet": {
                                                        "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
                                                    },
                                                },
                                            }
                                        ],
                                        "networkSecurityGroup": {
                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
                                        },
                                        "primary": True,
                                    },
                                }
                            ],
                        },
                        "osProfile": {
                            "adminPassword": "adfbrdxpv",
                            "adminUsername": "nrgzqciiaaxjrqldbmjbqkyhntp",
                            "allowExtensionOperations": True,
                            "computerNamePrefix": "o",
                            "customData": "xjjib",
                            "linuxConfiguration": {
                                "disablePasswordAuthentication": True,
                                "enableVMAgentPlatformUpdates": True,
                                "patchSettings": {
                                    "assessmentMode": "ImageDefault",
                                    "automaticByPlatformSettings": {
                                        "bypassPlatformSafetyChecksOnUserSchedule": True,
                                        "rebootSetting": "Unknown",
                                    },
                                    "patchMode": "ImageDefault",
                                },
                                "provisionVMAgent": True,
                                "ssh": {"publicKeys": [{"keyData": "kivgsubusvpprwqaqpjcmhsv", "path": "kmqz"}]},
                            },
                            "requireGuestProvisionSignal": True,
                            "secrets": [
                                {
                                    "sourceVault": {
                                        "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
                                    },
                                    "vaultCertificates": [
                                        {
                                            "certificateStore": "nlxrwavpzhueffxsshlun",
                                            "certificateUrl": "https://myVaultName.vault.azure.net/secrets/myCertName",
                                        }
                                    ],
                                }
                            ],
                            "windowsConfiguration": {
                                "additionalUnattendContent": [
                                    {
                                        "componentName": "Microsoft-Windows-Shell-Setup",
                                        "content": "bubmqbxjkj",
                                        "passName": "OobeSystem",
                                        "settingName": "AutoLogon",
                                    }
                                ],
                                "enableAutomaticUpdates": True,
                                "enableVMAgentPlatformUpdates": True,
                                "patchSettings": {
                                    "assessmentMode": "ImageDefault",
                                    "automaticByPlatformSettings": {
                                        "bypassPlatformSafetyChecksOnUserSchedule": True,
                                        "rebootSetting": "Unknown",
                                    },
                                    "enableHotpatching": True,
                                    "patchMode": "Manual",
                                },
                                "provisionVMAgent": True,
                                "timeZone": "hlyjiqcfksgrpjrct",
                                "winRM": {
                                    "listeners": [
                                        {
                                            "certificateUrl": "https://myVaultName.vault.azure.net/secrets/myCertName",
                                            "protocol": "Http",
                                        }
                                    ]
                                },
                            },
                        },
                        "scheduledEventsProfile": {
                            "osImageNotificationProfile": {
                                "enable": True,
                                "notBeforeTimeout": "olbpadmevekyczfokodtfprxti",
                            },
                            "terminateNotificationProfile": {"enable": True, "notBeforeTimeout": "iljppmmw"},
                        },
                        "securityPostureReference": {
                            "excludeExtensions": ["{securityPostureVMExtensionName}"],
                            "id": "/CommunityGalleries/{communityGalleryName}/securityPostures/{securityPostureName}/versions/{major.minor.patch}|{major.*}|latest",
                            "isOverridable": True,
                        },
                        "securityProfile": {
                            "encryptionAtHost": True,
                            "encryptionIdentity": {
                                "userAssignedIdentityResourceId": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{userAssignedIdentityName}"
                            },
                            "proxyAgentSettings": {"enabled": True, "keyIncarnationId": 20, "mode": "Audit"},
                            "securityType": "TrustedLaunch",
                            "uefiSettings": {"secureBootEnabled": True, "vTpmEnabled": True},
                        },
                        "serviceArtifactReference": {
                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/serviceArtifacts/{serviceArtifactsName}/vmArtifactsProfiles/{vmArtifactsProfileName}"
                        },
                        "storageProfile": {
                            "dataDisks": [
                                {
                                    "caching": "None",
                                    "createOption": "FromImage",
                                    "deleteOption": "Delete",
                                    "diskIOPSReadWrite": 27,
                                    "diskMBpsReadWrite": 2,
                                    "diskSizeGB": 6,
                                    "lun": 14,
                                    "managedDisk": {
                                        "diskEncryptionSet": {
                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
                                        },
                                        "securityProfile": {
                                            "diskEncryptionSet": {
                                                "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
                                            },
                                            "securityEncryptionType": "VMGuestStateOnly",
                                        },
                                        "storageAccountType": "Standard_LRS",
                                    },
                                    "name": "eogiykmdmeikswxmigjws",
                                    "writeAcceleratorEnabled": True,
                                }
                            ],
                            "diskControllerType": "uzb",
                            "imageReference": {
                                "communityGalleryImageId": "vlqe",
                                "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageName}/versions/{versionName}",
                                "offer": "isxgumkarlkomp",
                                "publisher": "mqxgwbiyjzmxavhbkd",
                                "sharedGalleryImageId": "kmkgihoxwlawuuhcinfirktdwkmx",
                                "sku": "eojmppqcrnpmxirtp",
                                "version": "wvpcqefgtmqdgltiuz",
                            },
                            "osDisk": {
                                "caching": "None",
                                "createOption": "FromImage",
                                "deleteOption": "Delete",
                                "diffDiskSettings": {"option": "Local", "placement": "CacheDisk"},
                                "diskSizeGB": 14,
                                "image": {
                                    "uri": "https://myStorageAccountName.blob.core.windows.net/myContainerName/myVhdName.vhd"
                                },
                                "managedDisk": {
                                    "diskEncryptionSet": {
                                        "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
                                    },
                                    "securityProfile": {
                                        "diskEncryptionSet": {
                                            "id": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
                                        },
                                        "securityEncryptionType": "VMGuestStateOnly",
                                    },
                                    "storageAccountType": "Standard_LRS",
                                },
                                "name": "wfttw",
                                "osType": "Windows",
                                "vhdContainers": ["tkzcwddtinkfpnfklatw"],
                                "writeAcceleratorEnabled": True,
                            },
                        },
                        "userData": "s",
                    },
                    "computeApiVersion": "2023-07-01",
                    "platformFaultDomainCount": 1,
                },
                "regularPriorityProfile": {"allocationStrategy": "LowestPrice", "capacity": 20, "minCapacity": 10},
                "spotPriorityProfile": {
                    "allocationStrategy": "PriceCapacityOptimized",
                    "capacity": 20,
                    "evictionPolicy": "Delete",
                    "maintain": True,
                    "maxPricePerVM": 0.00865,
                    "minCapacity": 10,
                },
                "vmSizesProfile": [{"name": "Standard_d1_v2", "rank": 19225}],
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-11-01/Fleets_Update.json
if __name__ == "__main__":
    main()
