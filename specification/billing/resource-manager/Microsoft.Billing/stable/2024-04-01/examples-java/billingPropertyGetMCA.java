
/**
 * Samples for BillingProperty Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyGetMCA.json
     */
    /**
     * Sample code: BillingPropertyGetMCA.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingPropertyGetMCA(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProperties().getWithResponse(null, null, com.azure.core.util.Context.NONE);
    }
}
