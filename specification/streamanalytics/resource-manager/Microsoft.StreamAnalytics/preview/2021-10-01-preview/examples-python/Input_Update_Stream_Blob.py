from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python input_update_stream_blob.py

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

    response = client.inputs.update(
        resource_group_name="sjrg8161",
        job_name="sj6695",
        input_name="input8899",
        input={
            "properties": {
                "datasource": {"properties": {"sourcePartitionCount": 32}, "type": "Microsoft.Storage/Blob"},
                "serialization": {"properties": {"encoding": "UTF8", "fieldDelimiter": "|"}, "type": "Csv"},
                "type": "Stream",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Update_Stream_Blob.json
if __name__ == "__main__":
    main()
