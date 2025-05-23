from azure.identity import DefaultAzureCredential

from azure.mgmt.devtestlabs import DevTestLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devtestlabs
# USAGE
    python formulas_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevTestLabsClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscriptionId}",
    )

    response = client.formulas.begin_create_or_update(
        resource_group_name="resourceGroupName",
        lab_name="{labName}",
        name="{formulaName}",
        formula={
            "location": "{location}",
            "properties": {
                "description": "Formula using a Linux base",
                "formulaContent": {
                    "location": "{location}",
                    "properties": {
                        "allowClaim": False,
                        "artifacts": [
                            {
                                "artifactId": "/artifactsources/{artifactSourceName}/artifacts/linux-install-nodejs",
                                "parameters": [],
                            }
                        ],
                        "disallowPublicIpAddress": True,
                        "galleryImageReference": {
                            "offer": "0001-com-ubuntu-server-groovy",
                            "osType": "Linux",
                            "publisher": "canonical",
                            "sku": "20_10",
                            "version": "latest",
                        },
                        "isAuthenticationWithSshKey": False,
                        "labSubnetName": "Dtl{labName}Subnet",
                        "labVirtualNetworkId": "/virtualnetworks/dtl{labName}",
                        "networkInterface": {
                            "sharedPublicIpAddressConfiguration": {
                                "inboundNatRules": [{"backendPort": 22, "transportProtocol": "Tcp"}]
                            }
                        },
                        "notes": "Ubuntu Server 20.10",
                        "size": "Standard_B1ms",
                        "storageType": "Standard",
                        "userName": "user",
                    },
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_CreateOrUpdate.json
if __name__ == "__main__":
    main()
