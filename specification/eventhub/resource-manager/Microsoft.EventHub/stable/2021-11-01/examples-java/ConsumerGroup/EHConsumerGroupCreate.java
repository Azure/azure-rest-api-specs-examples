import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.ConsumerGroupInner;

/** Samples for ConsumerGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/ConsumerGroup/EHConsumerGroupCreate.json
     */
    /**
     * Sample code: ConsumerGroupCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void consumerGroupCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getConsumerGroups()
            .createOrUpdateWithResponse(
                "ArunMonocle",
                "sdk-Namespace-2661",
                "sdk-EventHub-6681",
                "sdk-ConsumerGroup-5563",
                new ConsumerGroupInner().withUserMetadata("New consumergroup"),
                Context.NONE);
    }
}
