from azure.identity import DefaultAzureCredential
from azure.mgmt.timeseriesinsights import TimeSeriesInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-timeseriesinsights
# USAGE
    python reference_data_sets_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TimeSeriesInsightsClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.reference_data_sets.delete(
        resource_group_name="rg1",
        environment_name="env1",
        reference_data_set_name="rds1",
    )
    print(response)


# x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/ReferenceDataSetsDelete.json
if __name__ == "__main__":
    main()
