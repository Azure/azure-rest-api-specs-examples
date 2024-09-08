
/**
 * Samples for EnrollmentAccounts ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * enrollmentAccountsListByBillingAccount.json
     */
    /**
     * Sample code: EnrollmentAccountsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        enrollmentAccountsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().listByBillingAccount("6564892", null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
