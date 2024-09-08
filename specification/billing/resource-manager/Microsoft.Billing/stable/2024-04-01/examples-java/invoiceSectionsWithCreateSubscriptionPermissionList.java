
/**
 * Samples for BillingAccounts ListInvoiceSectionsByCreateSubscriptionPermission.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoiceSectionsWithCreateSubscriptionPermissionList.json
     */
    /**
     * Sample code: InvoiceSectionsWithCreateSubscriptionPermissionList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        invoiceSectionsWithCreateSubscriptionPermissionList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().listInvoiceSectionsByCreateSubscriptionPermission(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", null,
            com.azure.core.util.Context.NONE);
    }
}
