/** Samples for ReservationOrder Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationOrderDetailsWithExpandPlanInformation.json
     */
    /**
     * Sample code: GetReservationWithExpandPayments.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void getReservationWithExpandPayments(
        com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .reservationOrders()
            .getWithResponse("a075419f-44cc-497f-b68a-14ee811d48b9", "schedule", com.azure.core.util.Context.NONE);
    }
}
