
/**
 * Samples for SavingsPlanOrders ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * savingsPlanOrderListByBillingAccount.json
     */
    /**
     * Sample code: SavingsPlanOrderList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void savingsPlanOrderList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.savingsPlanOrders().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
