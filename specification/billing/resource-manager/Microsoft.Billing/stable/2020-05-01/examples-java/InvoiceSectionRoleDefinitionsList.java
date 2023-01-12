/** Samples for BillingRoleDefinitions ListByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleDefinitionsList.json
     */
    /**
     * Sample code: InvoiceSectionRoleDefinitionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionRoleDefinitionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleDefinitions()
            .listByInvoiceSection(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                com.azure.core.util.Context.NONE);
    }
}
