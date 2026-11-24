
import com.azure.resourcemanager.servicebus.fluent.models.SBSubscriptionInner;

/**
 * Samples for Subscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Subscriptions/SBSubscriptionCreate.json
     */
    /**
     * Sample code: SubscriptionCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void subscriptionCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getSubscriptions().createOrUpdateWithResponse("ResourceGroup", "sdk-Namespace-1349",
            "sdk-Topics-8740", "sdk-Subscriptions-2178", new SBSubscriptionInner().withEnableBatchedOperations(true),
            com.azure.core.util.Context.NONE);
    }
}
