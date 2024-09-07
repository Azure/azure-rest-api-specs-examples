
/**
 * Samples for Customers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/customersGet.json
     */
    /**
     * Sample code: CustomersGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void customersGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.customers().getWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
