
/**
 * Samples for ConsumerGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ConsumerGroup/EHConsumerGroupDelete.json
     */
    /**
     * Sample code: ConsumerGroupDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void consumerGroupDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getConsumerGroups().deleteWithResponse("ArunMonocle", "sdk-Namespace-2661",
            "sdk-EventHub-6681", "sdk-ConsumerGroup-5563", com.azure.core.util.Context.NONE);
    }
}
