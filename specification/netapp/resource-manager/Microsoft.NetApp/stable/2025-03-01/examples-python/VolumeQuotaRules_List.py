from azure.identity import DefaultAzureCredential

from azure.mgmt.netapp import NetAppManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-netapp
# USAGE
    python volume_quota_rules_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetAppManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="5275316f-a498-48d6-b324-2cbfdc4311b9",
    )

    response = client.volume_quota_rules.list_by_volume(
        resource_group_name="myRG",
        account_name="account-9957",
        pool_name="pool-5210",
        volume_name="volume-6387",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/VolumeQuotaRules_List.json
if __name__ == "__main__":
    main()
