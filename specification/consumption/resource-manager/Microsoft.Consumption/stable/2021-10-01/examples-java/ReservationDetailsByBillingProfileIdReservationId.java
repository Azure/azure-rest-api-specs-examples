
/**
 * Samples for ReservationsDetails List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationDetailsByBillingProfileIdReservationId.json
     */
    /**
     * Sample code: ReservationDetailsByBillingProfileIdReservationId.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationDetailsByBillingProfileIdReservationId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationsDetails().list(
            "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579", "2019-09-01", "2019-10-31",
            null, "1c6b6358-709f-484c-85f1-72e862a0cf3b", "9f39ba10-794f-4dcb-8f4b-8d0cb47c27dc",
            com.azure.core.util.Context.NONE);
    }
}
