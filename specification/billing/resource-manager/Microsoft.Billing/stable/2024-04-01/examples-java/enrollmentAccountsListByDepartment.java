
/**
 * Samples for EnrollmentAccounts ListByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * enrollmentAccountsListByDepartment.json
     */
    /**
     * Sample code: EnrollmentAccountsListByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void enrollmentAccountsListByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().listByDepartment("6564892", "164821", null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
