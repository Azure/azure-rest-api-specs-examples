
/**
 * Samples for Customers ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/customersListByBillingAccount
     * .json
     */
    /**
     * Sample code: CustomersListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void customersListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.customers().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", null, null, null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
