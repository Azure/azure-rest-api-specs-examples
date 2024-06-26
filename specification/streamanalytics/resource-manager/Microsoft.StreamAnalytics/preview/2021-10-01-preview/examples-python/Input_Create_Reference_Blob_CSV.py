from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python input_create_reference_blob_csv.py

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

    response = client.inputs.create_or_replace(
        resource_group_name="sjrg8440",
        job_name="sj9597",
        input_name="input7225",
        input={
            "properties": {
                "datasource": {
                    "properties": {
                        "blobName": "myblobinput",
                        "container": "state",
                        "dateFormat": "yyyy/MM/dd",
                        "deltaPathPattern": "/testBlob/{date}/delta/{time}/",
                        "deltaSnapshotRefreshRate": "16:14:30",
                        "fullSnapshotRefreshRate": "16:14:30",
                        "pathPattern": "{date}/{time}",
                        "sourcePartitionCount": 16,
                        "storageAccounts": [{"accountKey": "someAccountKey==", "accountName": "someAccountName"}],
                        "timeFormat": "HH",
                    },
                    "type": "Microsoft.Storage/Blob",
                },
                "serialization": {"properties": {"encoding": "UTF8", "fieldDelimiter": ","}, "type": "Csv"},
                "type": "Reference",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_Reference_Blob_CSV.json
if __name__ == "__main__":
    main()
