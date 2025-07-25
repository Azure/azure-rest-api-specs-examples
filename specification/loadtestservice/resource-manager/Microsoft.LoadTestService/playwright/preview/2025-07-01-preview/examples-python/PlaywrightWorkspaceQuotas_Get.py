from azure.identity import DefaultAzureCredential

from azure.mgmt.playwright import PlaywrightMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-playwright
# USAGE
    python playwright_workspace_quotas_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PlaywrightMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.playwright_workspace_quotas.get(
        resource_group_name="dummyrg",
        playwright_workspace_name="myWorkspace",
        quota_name="ExecutionMinutes",
    )
    print(response)


# x-ms-original-file: 2025-07-01-preview/PlaywrightWorkspaceQuotas_Get.json
if __name__ == "__main__":
    main()
