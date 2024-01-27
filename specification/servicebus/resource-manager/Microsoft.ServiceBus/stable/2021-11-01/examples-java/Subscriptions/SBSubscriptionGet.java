
import com.azure.core.util.Context;

/** Samples for Subscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/
     * SBSubscriptionGet.json
     */
    /**
     * Sample code: SubscriptionGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void subscriptionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getSubscriptions().getWithResponse("ResourceGroup",
            "sdk-Namespace-1349", "sdk-Topics-8740", "sdk-Subscriptions-2178", Context.NONE);
    }
}
