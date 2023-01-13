/** Samples for Invoices ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoicesListByBillingProfileWithRebillDetails.json
     */
    /**
     * Sample code: InvoicesListByBillingProfileWithRebillDetails.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesListByBillingProfileWithRebillDetails(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .invoices()
            .listByBillingProfile(
                "{billingAccountName}",
                "{billingProfileName}",
                "2018-01-01",
                "2018-06-30",
                com.azure.core.util.Context.NONE);
    }
}
