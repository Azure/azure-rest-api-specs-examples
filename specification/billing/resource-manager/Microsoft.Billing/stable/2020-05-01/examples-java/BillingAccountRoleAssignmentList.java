/** Samples for BillingRoleAssignments ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignmentList.json
     */
    /**
     * Sample code: BillingAccountRoleAssignmentList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountRoleAssignmentList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().listByBillingAccount("{billingAccountName}", com.azure.core.util.Context.NONE);
    }
}
