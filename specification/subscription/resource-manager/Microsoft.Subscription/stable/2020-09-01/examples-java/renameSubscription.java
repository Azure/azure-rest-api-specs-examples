
import com.azure.resourcemanager.subscription.models.SubscriptionName;

/**
 * Samples for SubscriptionOperation Rename.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2020-09-01/examples/renameSubscription.
     * json
     */
    /**
     * Sample code: renameSubscription.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void renameSubscription(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptionOperations().renameWithResponse("83aa47df-e3e9-49ff-877b-94304bf3d3ad",
            new SubscriptionName().withSubscriptionName("Test Sub"), com.azure.core.util.Context.NONE);
    }
}
