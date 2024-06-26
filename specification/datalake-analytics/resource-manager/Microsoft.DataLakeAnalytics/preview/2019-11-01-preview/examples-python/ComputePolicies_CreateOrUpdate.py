from azure.identity import DefaultAzureCredential
from azure.mgmt.datalake.analytics import DataLakeAnalyticsAccountManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datalake-analytics
# USAGE
    python compute_policies_create_or_update.py

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

    response = client.compute_policies.create_or_update(
        resource_group_name="contosorg",
        account_name="contosoadla",
        compute_policy_name="test_policy",
        parameters={
            "properties": {
                "maxDegreeOfParallelismPerJob": 10,
                "minPriorityPerJob": 30,
                "objectId": "776b9091-8916-4638-87f7-9c989a38da98",
                "objectType": "User",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_CreateOrUpdate.json
if __name__ == "__main__":
    main()
