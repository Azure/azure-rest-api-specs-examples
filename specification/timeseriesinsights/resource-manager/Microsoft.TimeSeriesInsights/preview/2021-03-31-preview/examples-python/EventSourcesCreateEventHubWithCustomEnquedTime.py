from azure.identity import DefaultAzureCredential
from azure.mgmt.timeseriesinsights import TimeSeriesInsightsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-timeseriesinsights
# USAGE
    python event_sources_create_event_hub_with_custom_enqued_time.py

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

    response = client.event_sources.create_or_update(
        resource_group_name="rg1",
        environment_name="env1",
        event_source_name="es1",
        parameters={
            "kind": "Microsoft.EventHub",
            "location": "West US",
            "properties": {
                "consumerGroupName": "cgn",
                "eventHubName": "ehn",
                "eventSourceResourceId": "somePathInArm",
                "ingressStartAt": {"time": "2017-04-01T19:20:33.2288820Z", "type": "CustomEnqueuedTime"},
                "keyName": "managementKey",
                "serviceBusNamespace": "sbn",
                "sharedAccessKey": "someSecretvalue",
                "timestampPropertyName": "someTimestampProperty",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EventSourcesCreateEventHubWithCustomEnquedTime.json
if __name__ == "__main__":
    main()
