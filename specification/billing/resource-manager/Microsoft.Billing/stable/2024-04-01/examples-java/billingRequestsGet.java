
/**
 * Samples for BillingRequests Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsGet.json
     */
    /**
     * Sample code: BillingRequestsGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRequestsGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRequests().getWithResponse("00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
