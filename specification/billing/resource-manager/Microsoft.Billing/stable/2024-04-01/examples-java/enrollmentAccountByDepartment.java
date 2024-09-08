
/**
 * Samples for EnrollmentAccounts GetByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountByDepartment
     * .json
     */
    /**
     * Sample code: EnrollmentAccountByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void enrollmentAccountByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().getByDepartmentWithResponse("6564892", "164821", "257698",
            com.azure.core.util.Context.NONE);
    }
}
