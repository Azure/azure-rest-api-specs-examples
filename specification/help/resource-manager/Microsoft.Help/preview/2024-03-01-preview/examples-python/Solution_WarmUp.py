from azure.identity import DefaultAzureCredential

from azure.mgmt.selfhelp import SelfHelpMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-selfhelp
# USAGE
    python solution_warm_up.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SelfHelpMgmtClient(
        credential=DefaultAzureCredential(),
    )

    client.solution.warm_up(
        scope="subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
        solution_resource_name="SolutionResourceName1",
    )


# x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Solution_WarmUp.json
if __name__ == "__main__":
    main()
