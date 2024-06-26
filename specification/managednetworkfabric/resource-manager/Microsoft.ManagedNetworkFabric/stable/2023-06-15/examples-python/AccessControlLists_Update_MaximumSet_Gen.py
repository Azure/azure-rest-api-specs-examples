from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python access_control_lists_update_maximum_set_gen.py

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

    response = client.access_control_lists.begin_update(
        resource_group_name="example-rg",
        access_control_list_name="example-acl",
        body={
            "properties": {
                "aclsUrl": "https://microsoft.com/a",
                "annotation": "annotation",
                "configurationType": "File",
                "dynamicMatchConfigurations": [
                    {
                        "ipGroups": [
                            {"ipAddressType": "IPv4", "ipPrefixes": ["10.20.3.1/20"], "name": "example-ipGroup"}
                        ],
                        "portGroups": [{"name": "example-portGroup", "ports": ["100-200"]}],
                        "vlanGroups": [{"name": "example-vlanGroup", "vlans": ["20-30"]}],
                    }
                ],
                "matchConfigurations": [
                    {
                        "actions": [{"counterName": "example-counter", "type": "Count"}],
                        "ipAddressType": "IPv4",
                        "matchConditions": [
                            {
                                "dscpMarkings": ["32"],
                                "etherTypes": ["0x1"],
                                "fragments": ["0xff00-0xffff"],
                                "ipCondition": {
                                    "ipGroupNames": ["example-ipGroup"],
                                    "ipPrefixValues": ["10.20.20.20/12"],
                                    "prefixType": "Prefix",
                                    "type": "SourceIP",
                                },
                                "ipLengths": ["4094-9214"],
                                "portCondition": {
                                    "flags": ["established"],
                                    "layer4Protocol": "TCP",
                                    "portGroupNames": ["example-portGroup"],
                                    "portType": "SourcePort",
                                    "ports": ["1-20"],
                                },
                                "protocolTypes": ["TCP"],
                                "ttlValues": ["23"],
                                "vlanMatchCondition": {
                                    "innerVlans": ["30"],
                                    "vlanGroupNames": ["example-vlanGroup"],
                                    "vlans": ["20-30"],
                                },
                            }
                        ],
                        "matchConfigurationName": "example-match",
                        "sequenceNumber": 123,
                    }
                ],
            },
            "tags": {"keyID": "KeyValue"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
