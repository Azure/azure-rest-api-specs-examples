
/**
 * Samples for Subscription AcceptOwnershipStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/
     * acceptOwnershipStatus.json
     */
    /**
     * Sample code: AcceptOwnershipStatus.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void acceptOwnershipStatus(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptions().acceptOwnershipStatusWithResponse("291bba3f-e0a5-47bc-a099-3bdcb2a50a05",
            com.azure.core.util.Context.NONE);
    }
}
