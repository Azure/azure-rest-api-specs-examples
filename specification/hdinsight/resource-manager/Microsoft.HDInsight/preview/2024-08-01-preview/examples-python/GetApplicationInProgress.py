from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsight import HDInsightManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsight
# USAGE
    python get_application_in_progress.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HDInsightManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.applications.get(
        resource_group_name="rg1",
        cluster_name="cluster1",
        application_name="app",
    )
    print(response)


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetApplicationInProgress.json
if __name__ == "__main__":
    main()
