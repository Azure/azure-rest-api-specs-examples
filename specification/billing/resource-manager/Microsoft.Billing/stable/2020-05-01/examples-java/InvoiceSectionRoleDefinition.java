/** Samples for BillingRoleDefinitions GetByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleDefinition.json
     */
    /**
     * Sample code: InvoiceSectionRoleDefinition.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionRoleDefinition(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleDefinitions()
            .getByInvoiceSectionWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{billingRoleDefinitionName}",
                com.azure.core.util.Context.NONE);
    }
}
