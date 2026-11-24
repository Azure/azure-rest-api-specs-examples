
/**
 * Samples for Subscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Subscriptions/SBSubscriptionGet.json
     */
    /**
     * Sample code: SubscriptionGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void subscriptionGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getSubscriptions().getWithResponse("ResourceGroup", "sdk-Namespace-1349",
            "sdk-Topics-8740", "sdk-Subscriptions-2178", com.azure.core.util.Context.NONE);
    }
}
