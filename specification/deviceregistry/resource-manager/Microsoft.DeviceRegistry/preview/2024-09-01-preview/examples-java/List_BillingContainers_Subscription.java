
/**
 * Samples for BillingContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_BillingContainers_Subscription.json
     */
    /**
     * Sample code: List_BillingContainers_Subscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listBillingContainersSubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.billingContainers().list(com.azure.core.util.Context.NONE);
    }
}
