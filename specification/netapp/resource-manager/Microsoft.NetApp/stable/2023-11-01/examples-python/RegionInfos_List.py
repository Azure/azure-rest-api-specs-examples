from azure.identity import DefaultAzureCredential

from azure.mgmt.netapp import NetAppManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-netapp
# USAGE
    python region_infos_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetAppManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="D633CC2E-722B-4AE1-B636-BBD9E4C60ED9",
    )

    response = client.net_app_resource_region_infos.list(
        location="eastus",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/RegionInfos_List.json
if __name__ == "__main__":
    main()
