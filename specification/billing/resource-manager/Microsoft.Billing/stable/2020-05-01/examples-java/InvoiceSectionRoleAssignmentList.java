import com.azure.core.util.Context;

/** Samples for BillingRoleAssignments ListByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignmentList.json
     */
    /**
     * Sample code: InvoiceSectionRoleAssignmentList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionRoleAssignmentList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .listByInvoiceSection("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", Context.NONE);
    }
}
