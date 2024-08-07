from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.informaticadatamanagement import InformaticaDataMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-informaticadatamanagement
# USAGE
    python serverless_runtimes_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = InformaticaDataMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="3599DA28-E346-4D9F-811E-189C0445F0FE",
    )

    response = client.serverless_runtimes.update(
        resource_group_name="rgopenapi",
        organization_name="W5",
        serverless_runtime_name="t_",
        properties={
            "properties": {
                "advancedCustomProperties": [{"key": "qcmc", "value": "unraxmnohdmvutt"}],
                "applicationType": "CDI",
                "computeUnits": "uncwbpu",
                "description": "ocprslpljoikxyduackzqnkuhyzrh",
                "executionTimeout": "tjyfytuywriabt",
                "platform": "AZURE",
                "serverlessAccountLocation": "goaugkyfanqfnvcmntreibqrswfpis",
                "serverlessRuntimeConfig": {
                    "cdiConfigProps": [
                        {
                            "applicationConfigs": [
                                {
                                    "customized": "j",
                                    "defaultValue": "zvgkqwmi",
                                    "name": "upfvjrqcrwwedfujkmsodeinw",
                                    "platform": "dixfyeobngivyvf",
                                    "type": "lw",
                                    "value": "mozgsetpwjmtyl",
                                }
                            ],
                            "engineName": "hngsdqvtjdhwqlbqfotipaiwjuys",
                            "engineVersion": "zlrlbg",
                        }
                    ],
                    "cdieConfigProps": [
                        {
                            "applicationConfigs": [
                                {
                                    "customized": "j",
                                    "defaultValue": "zvgkqwmi",
                                    "name": "upfvjrqcrwwedfujkmsodeinw",
                                    "platform": "dixfyeobngivyvf",
                                    "type": "lw",
                                    "value": "mozgsetpwjmtyl",
                                }
                            ],
                            "engineName": "hngsdqvtjdhwqlbqfotipaiwjuys",
                            "engineVersion": "zlrlbg",
                        }
                    ],
                },
                "serverlessRuntimeNetworkProfile": {
                    "networkInterfaceConfiguration": {
                        "subnetId": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/subnet1",
                        "vnetId": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1",
                        "vnetResourceGuid": "5328d299-1462-4be0-bef1-303a28e556a0",
                    }
                },
                "serverlessRuntimeTags": [{"name": "korveuycuwhs", "value": "uyiuegxnkgp"}],
                "serverlessRuntimeUserContextProperties": {"userContextToken": "ctgebtvjhwh"},
                "supplementaryFileLocation": "csxaqzpxu",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
