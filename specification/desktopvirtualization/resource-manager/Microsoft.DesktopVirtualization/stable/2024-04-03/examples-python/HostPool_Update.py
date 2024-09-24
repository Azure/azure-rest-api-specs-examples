from azure.identity import DefaultAzureCredential

from azure.mgmt.desktopvirtualization import DesktopVirtualizationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-desktopvirtualization
# USAGE
    python host_pool_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DesktopVirtualizationMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="daefabc0-95b4-48b3-b645-8a753a63c4fa",
    )

    response = client.host_pools.update(
        resource_group_name="resourceGroup1",
        host_pool_name="hostPool1",
    )
    print(response)


# x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/HostPool_Update.json
if __name__ == "__main__":
    main()
