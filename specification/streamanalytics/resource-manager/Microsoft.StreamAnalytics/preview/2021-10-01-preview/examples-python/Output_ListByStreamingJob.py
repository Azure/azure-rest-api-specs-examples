from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python list_all_outputs_in_a_streaming_job.py

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

    response = client.outputs.list_by_streaming_job(
        resource_group_name="sjrg2157",
        job_name="sj6458",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_ListByStreamingJob.json
if __name__ == "__main__":
    main()
