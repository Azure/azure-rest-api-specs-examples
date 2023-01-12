/** Samples for BillingPermissions ListByInvoiceSections. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionPermissionsList.json
     */
    /**
     * Sample code: InvoiceSectionPermissionsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionPermissionsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingPermissions()
            .listByInvoiceSections(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                com.azure.core.util.Context.NONE);
    }
}
