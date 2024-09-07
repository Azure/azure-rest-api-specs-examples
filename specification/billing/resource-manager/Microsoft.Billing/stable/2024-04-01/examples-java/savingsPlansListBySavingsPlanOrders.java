
/**
 * Samples for SavingsPlans ListBySavingsPlanOrder.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * savingsPlansListBySavingsPlanOrders.json
     */
    /**
     * Sample code: SavingsPlansInOrderList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void savingsPlansInOrderList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.savingsPlans().listBySavingsPlanOrder(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "20000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
