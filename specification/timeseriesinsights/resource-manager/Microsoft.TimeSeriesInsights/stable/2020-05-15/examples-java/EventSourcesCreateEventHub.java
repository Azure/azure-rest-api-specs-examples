import com.azure.resourcemanager.timeseriesinsights.models.EventHubEventSourceCreateOrUpdateParameters;
import com.azure.resourcemanager.timeseriesinsights.models.IngressStartAtType;
import com.azure.resourcemanager.timeseriesinsights.models.LocalTimestamp;
import com.azure.resourcemanager.timeseriesinsights.models.LocalTimestampFormat;
import com.azure.resourcemanager.timeseriesinsights.models.LocalTimestampTimeZoneOffset;

/** Samples for EventSources CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EventSourcesCreateEventHub.json
     */
    /**
     * Sample code: CreateEventHubEventSource.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void createEventHubEventSource(
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
                    .withLocalTimestampPropertiesLocalTimestamp(
                        new LocalTimestamp()
                            .withFormat(LocalTimestampFormat.fromString("TimeSpan"))
                            .withTimeZoneOffset(
                                new LocalTimestampTimeZoneOffset().withPropertyName("someEventPropertyName")))
                    .withType(IngressStartAtType.EARLIEST_AVAILABLE),
                com.azure.core.util.Context.NONE);
    }
}
