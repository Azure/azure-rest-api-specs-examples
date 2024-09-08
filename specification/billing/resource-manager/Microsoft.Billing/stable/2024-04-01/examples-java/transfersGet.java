
/**
 * Samples for Transfers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transfersGet.json
     */
    /**
     * Sample code: TransferGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void transferGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.transfers().getWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", "aabb123", com.azure.core.util.Context.NONE);
    }
}
