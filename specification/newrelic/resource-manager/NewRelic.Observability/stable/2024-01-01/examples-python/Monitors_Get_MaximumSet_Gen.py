from azure.identity import DefaultAzureCredential
from azure.mgmt.newrelicobservability import NewRelicObservabilityMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-newrelicobservability
# USAGE
    python monitors_get_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NewRelicObservabilityMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.monitors.get(
        resource_group_name="rgNewRelic",
        monitor_name="cdlymktqw",
    )
    print(response)


# x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/Monitors_Get_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
