
/**
 * Samples for Subscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Subscriptions/SBSubscriptionDelete.json
     */
    /**
     * Sample code: SubscriptionDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void subscriptionDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getSubscriptions().deleteWithResponse("ResourceGroup", "sdk-Namespace-5882",
            "sdk-Topics-1804", "sdk-Subscriptions-3670", com.azure.core.util.Context.NONE);
    }
}
