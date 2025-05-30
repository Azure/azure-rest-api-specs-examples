from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python nsp_logging_configuration_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId",
    )

    response = client.network_security_perimeter_logging_configurations.create_or_update(
        resource_group_name="rg1",
        network_security_perimeter_name="nsp1",
        logging_configuration_name="instance",
        parameters={
            "properties": {
                "enabledLogCategories": [
                    "NspPublicInboundPerimeterRulesDenied",
                    "NspPublicOutboundPerimeterRulesDenied",
                ]
            }
        },
    )
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspLoggingConfigurationPut.json
if __name__ == "__main__":
    main()
