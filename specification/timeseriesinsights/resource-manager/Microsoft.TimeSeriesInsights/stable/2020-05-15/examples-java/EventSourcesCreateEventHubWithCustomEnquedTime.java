import com.azure.resourcemanager.timeseriesinsights.models.EventHubEventSourceCreateOrUpdateParameters;
import com.azure.resourcemanager.timeseriesinsights.models.IngressStartAtType;

/** Samples for EventSources CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EventSourcesCreateEventHubWithCustomEnquedTime.json
     */
    /**
     * Sample code: EventSourcesCreateEventHubWithCustomEnquedTime.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void eventSourcesCreateEventHubWithCustomEnquedTime(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager
            .eventSources()
            .createOrUpdateWithResponse(
                "rg1",
                "env1",
                "es1",
                new EventHubEventSourceCreateOrUpdateParameters()
                    .withLocation("West US")
                    .withSharedAccessKey("fakeTokenPlaceholder")
                    .withServiceBusNamespace("sbn")
                    .withEventHubName("ehn")
                    .withConsumerGroupName("cgn")
                    .withKeyName("fakeTokenPlaceholder")
                    .withEventSourceResourceId("somePathInArm")
                    .withTimestampPropertyName("someTimestampProperty")
                    .withType(IngressStartAtType.CUSTOM_ENQUEUED_TIME)
                    .withTime("2017-04-01T19:20:33.2288820Z"),
                com.azure.core.util.Context.NONE);
    }
}
