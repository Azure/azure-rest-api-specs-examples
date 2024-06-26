from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python function_create_csharp.py

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

    response = client.functions.create_or_replace(
        resource_group_name="sjrg",
        job_name="sjName",
        function_name="function588",
        function={
            "properties": {
                "properties": {
                    "binding": {
                        "properties": {
                            "class": "ASAEdgeUDFDemo.Class1",
                            "dllPath": "ASAEdgeApplication2_CodeBehind",
                            "method": "SquareFunction",
                            "updateMode": "Static",
                        },
                        "type": "Microsoft.StreamAnalytics/CLRUdf",
                    },
                    "inputs": [{"dataType": "nvarchar(max)"}],
                    "output": {"dataType": "nvarchar(max)"},
                },
                "type": "Scalar",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Create_CSharp.json
if __name__ == "__main__":
    main()
