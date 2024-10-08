
/**
 * Samples for SavingsPlans GetByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * savingsPlanGetExpandRenewPropertiesByBillingAccount.json
     */
    /**
     * Sample code: SavingsPlanGetExpandRenewProperties.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void savingsPlanGetExpandRenewProperties(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.savingsPlans().getByBillingAccountWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", null,
            com.azure.core.util.Context.NONE);
    }
}
