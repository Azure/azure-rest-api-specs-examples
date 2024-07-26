
/**
 * Samples for ConsumerGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ConsumerGroup/
     * EHConsumerGroupGet.json
     */
    /**
     * Sample code: ConsumerGroupGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void consumerGroupGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getConsumerGroups().getWithResponse("ArunMonocle",
            "sdk-Namespace-2661", "sdk-EventHub-6681", "sdk-ConsumerGroup-5563", com.azure.core.util.Context.NONE);
    }
}
