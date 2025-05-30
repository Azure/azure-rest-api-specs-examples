from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_network_mappings_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="9112a37f-0f3e-46ec-9c00-060c6edca071",
        resource_group_name="srcBvte2a14C27",
        resource_name="srce2avaultbvtaC27",
    )

    response = client.replication_network_mappings.begin_update(
        fabric_name="b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac",
        network_name="e2267b5c-2650-49bd-ab3f-d66aae694c06",
        network_mapping_name="corpe2amap",
        input={
            "properties": {
                "fabricSpecificDetails": {"instanceType": "VmmToAzure"},
                "recoveryFabricName": "Microsoft Azure",
                "recoveryNetworkId": "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai2",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationNetworkMappings_Update.json
if __name__ == "__main__":
    main()
