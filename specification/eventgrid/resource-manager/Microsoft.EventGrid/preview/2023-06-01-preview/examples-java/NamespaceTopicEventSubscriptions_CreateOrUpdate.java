import com.azure.resourcemanager.eventgrid.models.DeliveryConfiguration;
import com.azure.resourcemanager.eventgrid.models.DeliveryMode;
import com.azure.resourcemanager.eventgrid.models.DeliverySchema;
import com.azure.resourcemanager.eventgrid.models.QueueInfo;
import java.time.Duration;

/** Samples for NamespaceTopicEventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopicEventSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: NamespaceTopicEventSubscriptions_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicEventSubscriptionsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .namespaceTopicEventSubscriptions()
            .define("examplenamespacetopicEventSub2")
            .withExistingTopic("examplerg", "examplenamespace2", "examplenamespacetopic2")
            .withDeliveryConfiguration(
                new DeliveryConfiguration()
                    .withDeliveryMode(DeliveryMode.QUEUE)
                    .withQueue(
                        new QueueInfo()
                            .withReceiveLockDurationInSeconds(60)
                            .withMaxDeliveryCount(4)
                            .withEventTimeToLive(Duration.parse("P1D"))))
            .withEventDeliverySchema(DeliverySchema.CLOUD_EVENT_SCHEMA_V1_0)
            .create();
    }
}
