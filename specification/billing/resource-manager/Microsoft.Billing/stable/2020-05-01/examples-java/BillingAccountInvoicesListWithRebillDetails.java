import com.azure.core.util.Context;

/** Samples for Invoices ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountInvoicesListWithRebillDetails.json
     */
    /**
     * Sample code: BillingAccountInvoicesListWithRebillDetails.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountInvoicesListWithRebillDetails(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().listByBillingAccount("{billingAccountName}", "2018-01-01", "2018-06-30", Context.NONE);
    }
}
