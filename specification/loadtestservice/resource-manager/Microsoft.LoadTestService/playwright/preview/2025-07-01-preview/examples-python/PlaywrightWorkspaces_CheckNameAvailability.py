from azure.identity import DefaultAzureCredential

from azure.mgmt.playwright import PlaywrightMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-playwright
# USAGE
    python playwright_workspaces_check_name_availability.py

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

    response = client.playwright_workspaces.check_name_availability(
        body={"name": "dummyName", "type": "Microsoft.LoadTestService/PlaywrightWorkspaces"},
    )
    print(response)


# x-ms-original-file: 2025-07-01-preview/PlaywrightWorkspaces_CheckNameAvailability.json
if __name__ == "__main__":
    main()
