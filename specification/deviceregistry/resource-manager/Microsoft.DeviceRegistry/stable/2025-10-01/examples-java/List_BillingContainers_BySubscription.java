
/**
 * Samples for BillingContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_BillingContainers_BySubscription.json
     */
    /**
     * Sample code: List_BillingContainers_BySubscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listBillingContainersBySubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.billingContainers().list(com.azure.core.util.Context.NONE);
    }
}
