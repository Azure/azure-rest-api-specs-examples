import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.fluent.models.SBSubscriptionInner;

/** Samples for Subscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/SBSubscriptionCreate.json
     */
    /**
     * Sample code: SubscriptionCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void subscriptionCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .serviceBusNamespaces()
            .manager()
            .serviceClient()
            .getSubscriptions()
            .createOrUpdateWithResponse(
                "ResourceGroup",
                "sdk-Namespace-1349",
                "sdk-Topics-8740",
                "sdk-Subscriptions-2178",
                new SBSubscriptionInner().withEnableBatchedOperations(true),
                Context.NONE);
    }
}
