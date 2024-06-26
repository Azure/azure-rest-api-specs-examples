from azure.identity import DefaultAzureCredential
from azure.mgmt.datalake.analytics import DataLakeAnalyticsAccountManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datalake-analytics
# USAGE
    python accounts_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataLakeAnalyticsAccountManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.accounts.begin_create(
        resource_group_name="contosorg",
        account_name="contosoadla",
        parameters={
            "location": "eastus2",
            "properties": {
                "computePolicies": [
                    {
                        "name": "test_policy",
                        "properties": {
                            "maxDegreeOfParallelismPerJob": 1,
                            "minPriorityPerJob": 1,
                            "objectId": "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
                            "objectType": "User",
                        },
                    }
                ],
                "dataLakeStoreAccounts": [{"name": "test_adls", "properties": {"suffix": "test_suffix"}}],
                "defaultDataLakeStoreAccount": "test_adls",
                "firewallAllowAzureIps": "Enabled",
                "firewallRules": [
                    {"name": "test_rule", "properties": {"endIpAddress": "2.2.2.2", "startIpAddress": "1.1.1.1"}}
                ],
                "firewallState": "Enabled",
                "maxDegreeOfParallelism": 30,
                "maxDegreeOfParallelismPerJob": 1,
                "maxJobCount": 3,
                "minPriorityPerJob": 1,
                "newTier": "Consumption",
                "queryStoreRetention": 30,
                "storageAccounts": [
                    {
                        "name": "test_storage",
                        "properties": {"accessKey": "34adfa4f-cedf-4dc0-ba29-b6d1a69ab346", "suffix": "test_suffix"},
                    }
                ],
            },
            "tags": {"test_key": "test_value"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Create.json
if __name__ == "__main__":
    main()
