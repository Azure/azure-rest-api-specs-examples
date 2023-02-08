/** Samples for ReservationOrder List. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationOrders.json
     */
    /**
     * Sample code: ReservationOrderList.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void reservationOrderList(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.reservationOrders().list(com.azure.core.util.Context.NONE);
    }
}
