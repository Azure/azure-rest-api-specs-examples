
/**
 * Samples for Departments ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * departmentsListByBillingAccount.json
     */
    /**
     * Sample code: DepartmentsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void departmentsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.departments().listByBillingAccount("456598", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
