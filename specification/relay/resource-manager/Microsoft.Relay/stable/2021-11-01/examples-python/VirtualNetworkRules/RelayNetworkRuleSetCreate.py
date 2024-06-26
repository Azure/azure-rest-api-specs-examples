from azure.identity import DefaultAzureCredential
from azure.mgmt.relay import RelayAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-relay
# USAGE
    python name_space_network_rule_set_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RelayAPI(
        credential=DefaultAzureCredential(),
        subscription_id="Subscription",
    )

    response = client.namespaces.create_or_update_network_rule_set(
        resource_group_name="ResourceGroup",
        namespace_name="example-RelayNamespace-6019",
        parameters={
            "properties": {
                "defaultAction": "Deny",
                "ipRules": [
                    {"action": "Allow", "ipMask": "1.1.1.1"},
                    {"action": "Allow", "ipMask": "1.1.1.2"},
                    {"action": "Allow", "ipMask": "1.1.1.3"},
                    {"action": "Allow", "ipMask": "1.1.1.4"},
                    {"action": "Allow", "ipMask": "1.1.1.5"},
                ],
            }
        },
    )
    print(response)


# x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/VirtualNetworkRules/RelayNetworkRuleSetCreate.json
if __name__ == "__main__":
    main()
