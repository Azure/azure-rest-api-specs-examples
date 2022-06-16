import com.azure.core.util.Context;

/** Samples for BillingRoleAssignments GetByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignment.json
     */
    /**
     * Sample code: InvoiceSectionRoleAssignment.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionRoleAssignment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .getByInvoiceSectionWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{billingRoleAssignmentName}",
                Context.NONE);
    }
}
