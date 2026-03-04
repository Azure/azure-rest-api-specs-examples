
/**
 * Samples for BillingInfo Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/BillingInfo_Get.json
     */
    /**
     * Sample code: BillingInfo_Get.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        billingInfoGet(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.billingInfoes().getWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
