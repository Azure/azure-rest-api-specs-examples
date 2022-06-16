import com.azure.core.util.Context;

/** Samples for Invoices ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoicesListByBillingProfile.json
     */
    /**
     * Sample code: InvoicesListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .invoices()
            .listByBillingProfile(
                "{billingAccountName}", "{billingProfileName}", "2018-01-01", "2018-06-30", Context.NONE);
    }
}
