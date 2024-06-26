from azure.identity import DefaultAzureCredential
from azure.mgmt.testbase import TestBase

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-testbase
# USAGE
    python billing_hub_get_usage.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TestBase(
        credential=DefaultAzureCredential(),
        subscription_id="7b053408-0960-40b0-b950-99b0461aa9a8",
    )

    response = client.billing_hub_service.get_usage(
        resource_group_name="contoso-rg1",
        test_base_account_name="contoso-testBaseAccount1",
    )
    print(response)


# x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2022-04-01-preview/examples/BillingHubGetUsage.json
if __name__ == "__main__":
    main()
