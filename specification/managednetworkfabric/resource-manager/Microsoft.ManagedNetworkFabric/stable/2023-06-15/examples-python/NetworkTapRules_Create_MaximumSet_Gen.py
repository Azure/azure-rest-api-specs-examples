from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python network_tap_rules_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedNetworkFabricMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="1234ABCD-0A1B-1234-5678-123456ABCDEF",
    )

    response = client.network_tap_rules.begin_create(
        resource_group_name="example-rg",
        network_tap_rule_name="example-tapRule",
        body={
            "location": "eastus",
            "properties": {
                "annotation": "annotation",
                "configurationType": "File",
                "dynamicMatchConfigurations": [
                    {
                        "ipGroups": [
                            {"ipAddressType": "IPv4", "ipPrefixes": ["10.10.10.10/30"], "name": "example-ipGroup1"}
                        ],
                        "portGroups": [
                            {"name": "example-portGroup1", "ports": ["100-200"]},
                            {"name": "example-portGroup2", "ports": ["900", "1000-2000"]},
                        ],
                        "vlanGroups": [{"name": "exmaple-vlanGroup", "vlans": ["10", "100-200"]}],
                    }
                ],
                "matchConfigurations": [
                    {
                        "actions": [
                            {
                                "destinationId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup",
                                "isTimestampEnabled": "True",
                                "matchConfigurationName": "match1",
                                "truncate": "100",
                                "type": "Drop",
                            }
                        ],
                        "ipAddressType": "IPv4",
                        "matchConditions": [
                            {
                                "encapsulationType": "None",
                                "ipCondition": {
                                    "ipGroupNames": ["example-ipGroup"],
                                    "ipPrefixValues": ["10.10.10.10/20"],
                                    "prefixType": "Prefix",
                                    "type": "SourceIP",
                                },
                                "portCondition": {
                                    "layer4Protocol": "TCP",
                                    "portGroupNames": ["example-portGroup1"],
                                    "portType": "SourcePort",
                                    "ports": ["100"],
                                },
                                "protocolTypes": ["TCP"],
                                "vlanMatchCondition": {
                                    "innerVlans": ["11-20"],
                                    "vlanGroupNames": ["exmaple-vlanGroup"],
                                    "vlans": ["10"],
                                },
                            }
                        ],
                        "matchConfigurationName": "config1",
                        "sequenceNumber": 10,
                    }
                ],
                "pollingIntervalInSeconds": 30,
                "tapRulesUrl": "https://microsoft.com/a",
            },
            "tags": {"keyID": "keyValue"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTapRules_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
