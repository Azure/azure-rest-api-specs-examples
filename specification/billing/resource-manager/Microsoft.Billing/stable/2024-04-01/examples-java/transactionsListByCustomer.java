
import com.azure.resourcemanager.billing.models.TransactionType;
import java.time.LocalDate;

/**
 * Samples for Transactions ListByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionsListByCustomer.
     * json
     */
    /**
     * Sample code: TransactionsListByCustomer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void transactionsListByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.transactions().listByCustomer(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "22000000-0000-0000-0000-000000000000", LocalDate.parse("2024-04-01"), LocalDate.parse("2023-05-30"),
            TransactionType.BILLED, "properties/date gt '2020-10-01'", "properties/date", null, null, null, "storage",
            com.azure.core.util.Context.NONE);
    }
}
