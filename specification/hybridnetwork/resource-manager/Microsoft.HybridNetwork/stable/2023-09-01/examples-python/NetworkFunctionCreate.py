from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python network_function_create.py

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

    response = client.network_functions.begin_create_or_update(
        resource_group_name="rg",
        network_function_name="testNf",
        parameters={
            "location": "eastus",
            "properties": {
                "allowSoftwareUpdate": False,
                "configurationType": "Open",
                "deploymentValues": '{"releaseName":"testReleaseName","namespace":"testNamespace"}',
                "networkFunctionDefinitionVersionResourceReference": {
                    "id": "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1",
                    "idType": "Open",
                },
                "nfviId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation",
                "nfviType": "AzureArcKubernetes",
                "roleOverrideValues": [
                    '{"name":"testRoleOne","deployParametersMappingRuleProfile":{"helmMappingRuleProfile":{"helmPackageVersion":"2.1.3","values":"{\\"roleOneParam\\":\\"roleOneOverrideValue\\"}"}}}',
                    '{"name":"testRoleTwo","deployParametersMappingRuleProfile":{"helmMappingRuleProfile":{"releaseName":"overrideReleaseName","releaseNamespace":"overrideNamespace","values":"{\\"roleTwoParam\\":\\"roleTwoOverrideValue\\"}"}}}',
                ],
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionCreate.json
if __name__ == "__main__":
    main()
