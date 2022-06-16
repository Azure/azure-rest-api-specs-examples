import com.azure.core.util.Context;

/** Samples for BillingRoleAssignments DeleteByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignmentDelete.json
     */
    /**
     * Sample code: BillingAccountRoleAssignmentDelete.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountRoleAssignmentDelete(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .deleteByBillingAccountWithResponse("{billingAccountName}", "{billingRoleAssignmentName}", Context.NONE);
    }
}
