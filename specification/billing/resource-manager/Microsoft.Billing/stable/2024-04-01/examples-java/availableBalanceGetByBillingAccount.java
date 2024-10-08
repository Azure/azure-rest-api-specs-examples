
/**
 * Samples for AvailableBalances GetByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * availableBalanceGetByBillingAccount.json
     */
    /**
     * Sample code: AvailableBalanceGetByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void availableBalanceGetByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.availableBalances().getByBillingAccountWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            com.azure.core.util.Context.NONE);
    }
}
