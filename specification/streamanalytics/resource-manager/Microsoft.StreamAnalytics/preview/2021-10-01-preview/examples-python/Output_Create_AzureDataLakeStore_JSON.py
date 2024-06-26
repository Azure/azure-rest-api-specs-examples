from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python output_create_azure_data_lake_store_json.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StreamAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="56b5e0a9-b645-407d-99b0-c64f86013e3d",
    )

    response = client.outputs.create_or_replace(
        resource_group_name="sjrg6912",
        job_name="sj3310",
        output_name="output5195",
        output={
            "properties": {
                "datasource": {
                    "properties": {
                        "accountName": "someaccount",
                        "dateFormat": "yyyy/MM/dd",
                        "filePathPrefix": "{date}/{time}",
                        "refreshToken": "someRefreshToken==",
                        "tenantId": "cea4e98b-c798-49e7-8c40-4a2b3beb47dd",
                        "timeFormat": "HH",
                        "tokenUserDisplayName": "Bob Smith",
                        "tokenUserPrincipalName": "bobsmith@contoso.com",
                    },
                    "type": "Microsoft.DataLake/Accounts",
                },
                "serialization": {"properties": {"encoding": "UTF8", "format": "Array"}, "type": "Json"},
            }
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Create_AzureDataLakeStore_JSON.json
if __name__ == "__main__":
    main()
