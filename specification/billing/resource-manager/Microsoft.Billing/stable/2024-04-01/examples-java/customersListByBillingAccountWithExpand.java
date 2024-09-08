
/**
 * Samples for Customers ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * customersListByBillingAccountWithExpand.json
     */
    /**
     * Sample code: CustomersListByBillingAccountWithExpand.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        customersListByBillingAccountWithExpand(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.customers().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "enabledAzurePlans,resellers", null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
