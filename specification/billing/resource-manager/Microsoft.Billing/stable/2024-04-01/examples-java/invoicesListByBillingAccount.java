
import java.time.LocalDate;

/**
 * Samples for Invoices ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesListByBillingAccount.
     * json
     */
    /**
     * Sample code: InvoicesListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            LocalDate.parse("2023-01-01"), LocalDate.parse("2023-06-30"), null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
