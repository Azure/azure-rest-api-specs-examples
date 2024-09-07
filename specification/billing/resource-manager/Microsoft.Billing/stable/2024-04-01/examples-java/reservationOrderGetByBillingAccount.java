
/**
 * Samples for ReservationOrders GetByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * reservationOrderGetByBillingAccount.json
     */
    /**
     * Sample code: reservationOrderGetByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void reservationOrderGetByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.reservationOrders().getByBillingAccountWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "20000000-0000-0000-0000-000000000000", null, com.azure.core.util.Context.NONE);
    }
}
