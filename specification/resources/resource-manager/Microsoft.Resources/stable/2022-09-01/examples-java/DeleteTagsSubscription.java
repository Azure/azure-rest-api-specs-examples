
/** Samples for TagOperations DeleteAtScope. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/DeleteTagsSubscription.
     * json
     */
    /**
     * Sample code: Update tags on a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTagsOnASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getTagOperations()
            .deleteAtScope("subscriptions/00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
