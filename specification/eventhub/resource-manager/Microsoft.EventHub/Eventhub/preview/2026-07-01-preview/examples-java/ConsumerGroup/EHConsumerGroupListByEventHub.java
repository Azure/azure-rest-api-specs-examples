
/**
 * Samples for ConsumerGroups ListByEventHub.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ConsumerGroup/EHConsumerGroupListByEventHub.json
     */
    /**
     * Sample code: ConsumerGroupsListAll.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void consumerGroupsListAll(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getConsumerGroups().listByEventHub("ArunMonocle", "sdk-Namespace-2661",
            "sdk-EventHub-6681", null, null, com.azure.core.util.Context.NONE);
    }
}
