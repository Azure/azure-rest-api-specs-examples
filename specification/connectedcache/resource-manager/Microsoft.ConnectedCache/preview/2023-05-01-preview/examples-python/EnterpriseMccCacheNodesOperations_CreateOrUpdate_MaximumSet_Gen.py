from azure.identity import DefaultAzureCredential

from azure.mgmt.connectedcache import ConnectedCacheMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-connectedcache
# USAGE
    python enterprise_mcc_cache_nodes_operations_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConnectedCacheMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.enterprise_mcc_cache_nodes_operations.begin_create_or_update(
        resource_group_name="rgConnectedCache",
        customer_resource_name="nhdkvstdrrtsxxuz",
        cache_node_resource_name="fgduqdovidpemlnmhelomffuafdrbgaasqznvrdkbvspfzsnrhncdtqquhijhdpwyzwleukqldpceyxqhqlftqrr",
        resource={
            "location": "westus",
            "properties": {
                "additionalCacheNodeProperties": {
                    "bgpConfiguration": {"asnToIpAddressMapping": "fjbggfvumrn"},
                    "cacheNodePropertiesDetailsIssuesList": ["ennbzfpuszgalzpawmwicaofqcwcj"],
                    "driveConfiguration": [
                        {
                            "cacheNumber": 11,
                            "nginxMapping": "cirlpkpuxg",
                            "physicalPath": "pcbkezoofhamkycot",
                            "sizeInGb": 14,
                        }
                    ],
                    "optionalProperty1": "ph",
                    "optionalProperty2": "soqqgpgcbhb",
                    "optionalProperty3": "fpnycrbagptsujiotnjfuhlm",
                    "optionalProperty4": "gesqugrxvhxlxxyvatgrautxwlmxsf",
                    "optionalProperty5": "zknjgzpaqtvdqjydd",
                    "proxyUrl": "ihkzxlzvpcywtzrogupqozkdud",
                    "proxyUrlConfiguration": {"proxyUrl": "hplstyg"},
                },
                "cacheNode": {
                    "cacheNodeId": "fmrjefyddfn",
                    "cacheNodeName": "qppvqxliajjfoalzjbgmxggr",
                    "cidrCsv": ["nlqlvrthafvvljuupcbcw"],
                    "cidrSelectionType": 11,
                    "customerAsn": 25,
                    "customerIndex": "vafvezmelfgmjsrccjukrhppljvipg",
                    "customerName": "zsklcocjfjhkcpmzyefzkwamdzc",
                    "fullyQualifiedResourceId": "yeinlleavzbehelhsucb",
                    "ipAddress": "gbfkdhloyphnpnhemwrcrxlk",
                    "isEnabled": True,
                    "isEnterpriseManaged": True,
                    "maxAllowableEgressInMbps": 27,
                    "shouldMigrate": True,
                },
                "error": {},
                "statusCode": "1",
                "statusDetails": "lgljxmyyoakw",
                "statusText": "Success",
            },
            "tags": {"key259": "qbkixjuyjkoj"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2023-05-01-preview/EnterpriseMccCacheNodesOperations_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
