
/**
 * Samples for Products ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/productsListByBillingAccount.
     * json
     */
    /**
     * Sample code: ProductsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void productsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.products().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", null, null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
