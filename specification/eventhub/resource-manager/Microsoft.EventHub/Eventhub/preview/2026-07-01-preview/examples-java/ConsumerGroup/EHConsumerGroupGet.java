
/**
 * Samples for ConsumerGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ConsumerGroup/EHConsumerGroupGet.json
     */
    /**
     * Sample code: ConsumerGroupGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void consumerGroupGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getConsumerGroups().getWithResponse("ArunMonocle", "sdk-Namespace-2661",
            "sdk-EventHub-6681", "sdk-ConsumerGroup-5563", com.azure.core.util.Context.NONE);
    }
}
