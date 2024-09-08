
/**
 * Samples for Reservations ListByReservationOrder.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * reservationsGetFromOrderByBillingAccount.json
     */
    /**
     * Sample code: reservationsGetFromOrderByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        reservationsGetFromOrderByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.reservations().listByReservationOrder(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "20000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
