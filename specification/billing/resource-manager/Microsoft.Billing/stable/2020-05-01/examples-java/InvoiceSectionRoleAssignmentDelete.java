/** Samples for BillingRoleAssignments DeleteByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignmentDelete.json
     */
    /**
     * Sample code: InvoiceSectionRoleAssignmentDelete.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionRoleAssignmentDelete(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .deleteByInvoiceSectionWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{billingRoleAssignmentName}",
                com.azure.core.util.Context.NONE);
    }
}
