from azure.identity import DefaultAzureCredential

from azure.mgmt.scvmm import ScVmmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python vmm_servers_list_by_subscription_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ScVmmMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="79332E5A-630B-480F-A266-A941C015AB19",
    )

    response = client.vmm_servers.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_ListBySubscription_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
