from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python network_service_design_version_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.network_service_design_versions.begin_create_or_update(
        resource_group_name="rg",
        publisher_name="TestPublisher",
        network_service_design_group_name="TestNetworkServiceDesignGroupName",
        network_service_design_version_name="1.0.0",
        parameters={
            "location": "eastus",
            "properties": {
                "configurationGroupSchemaReferences": {
                    "MyVM_Configuration": {
                        "id": "/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/publishers/contosoGroup/networkServiceDesignGroups/NSD_contoso/configurationGroupSchemas/MyVM_Configuration_Schema"
                    }
                },
                "resourceElementTemplates": [
                    {
                        "configuration": {
                            "artifactProfile": {
                                "artifactName": "MyVMArmTemplate",
                                "artifactStoreReference": {
                                    "id": "/subscriptions/subid/providers/Microsoft.HybridNetwork/publishers/contosoGroup/artifactStoreReference/store1"
                                },
                                "artifactVersion": "1.0.0",
                            },
                            "parameterValues": '{"publisherName":"{configurationparameters(\'MyVM_Configuration\').publisherName}","skuGroupName":"{configurationparameters(\'MyVM_Configuration\').skuGroupName}","skuVersion":"{configurationparameters(\'MyVM_Configuration\').skuVersion}","skuOfferingLocation":"{configurationparameters(\'MyVM_Configuration\').skuOfferingLocation}","nfviType":"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviType}","nfviId":"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviId}","allowSoftwareUpdates":"{configurationparameters(\'MyVM_Configuration\').allowSoftwareUpdates}","virtualNetworkName":"{configurationparameters(\'MyVM_Configuration\').vnetName}","subnetName":"{configurationparameters(\'MyVM_Configuration\').subnetName}","subnetAddressPrefix":"{configurationparameters(\'MyVM_Configuration\').subnetAddressPrefix}","managedResourceGroup":"{configurationparameters(\'SNSSelf\').managedResourceGroupName}","adminPassword":"{secretparameters(\'MyVM_Configuration\').adminPassword}"}',
                            "templateType": "ArmTemplate",
                        },
                        "dependsOnProfile": {"installDependsOn": []},
                        "name": "MyVM",
                        "type": "ArmResourceDefinition",
                    }
                ],
                "versionState": "Active",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkServiceDesignVersionCreate.json
if __name__ == "__main__":
    main()
