from azure.identity import DefaultAzureCredential

from azure.mgmt.dataprotection import DataProtectionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dataprotection
# USAGE
    python get_resource_guard_proxy.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataProtectionMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="5e13b949-1218-4d18-8b99-7e12155ec4f7",
    )

    response = client.dpp_resource_guard_proxy.get(
        resource_group_name="SampleResourceGroup",
        vault_name="sampleVault",
        resource_guard_proxy_name="swaggerExample",
    )
    print(response)


# x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/ResourceGuardProxyCRUD/GetResourceGuardProxy.json
if __name__ == "__main__":
    main()
