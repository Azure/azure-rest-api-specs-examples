import com.azure.core.util.Context;

/** Samples for BillingAccounts ListInvoiceSectionsByCreateSubscriptionPermission. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionsListWithCreateSubPermission.json
     */
    /**
     * Sample code: InvoiceSectionsListWithCreateSubPermission.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionsListWithCreateSubPermission(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingAccounts()
            .listInvoiceSectionsByCreateSubscriptionPermission("{billingAccountName}", Context.NONE);
    }
}
