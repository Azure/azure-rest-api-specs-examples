from azure.identity import DefaultAzureCredential

from azure.mgmt.quota import QuotaMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-quota
# USAGE
    python get_quota_requests_history.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = QuotaMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.quota_request_status.list(
        scope="subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/getQuotaRequestsHistory.json
if __name__ == "__main__":
    main()
