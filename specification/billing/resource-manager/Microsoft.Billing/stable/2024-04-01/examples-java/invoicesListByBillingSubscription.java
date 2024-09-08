
import java.time.LocalDate;

/**
 * Samples for Invoices ListByBillingSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesListByBillingSubscription.json
     */
    /**
     * Sample code: InvoicesListByBillingSubscription.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesListByBillingSubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().listByBillingSubscription(LocalDate.parse("2023-01-01"), LocalDate.parse("2023-06-30"), null,
            null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
