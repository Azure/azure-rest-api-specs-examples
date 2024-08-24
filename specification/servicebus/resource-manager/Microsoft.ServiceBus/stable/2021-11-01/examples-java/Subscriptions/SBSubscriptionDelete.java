
/**
 * Samples for Subscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/
     * SBSubscriptionDelete.json
     */
    /**
     * Sample code: SubscriptionDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void subscriptionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getSubscriptions().deleteWithResponse("ResourceGroup",
            "sdk-Namespace-5882", "sdk-Topics-1804", "sdk-Subscriptions-3670", com.azure.core.util.Context.NONE);
    }
}
