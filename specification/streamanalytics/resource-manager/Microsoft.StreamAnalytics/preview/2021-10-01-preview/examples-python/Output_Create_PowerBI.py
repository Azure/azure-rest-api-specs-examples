from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python output_create_power_bi.py

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
        resource_group_name="sjrg7983",
        job_name="sj2331",
        output_name="output3022",
        output={
            "properties": {
                "datasource": {
                    "properties": {
                        "dataset": "someDataset",
                        "groupId": "ac40305e-3e8d-43ac-8161-c33799f43e95",
                        "groupName": "MyPowerBIGroup",
                        "refreshToken": "someRefreshToken==",
                        "table": "someTable",
                        "tokenUserDisplayName": "Bob Smith",
                        "tokenUserPrincipalName": "bobsmith@contoso.com",
                    },
                    "type": "PowerBI",
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Create_PowerBI.json
if __name__ == "__main__":
    main()
