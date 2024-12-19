
/**
 * Samples for BillingInfo Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/BillingInfo_Get.json
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
