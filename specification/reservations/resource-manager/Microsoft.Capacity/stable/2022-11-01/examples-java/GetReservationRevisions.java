/** Samples for Reservation ListRevisions. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationRevisions.json
     */
    /**
     * Sample code: ReservationRevisions.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void reservationRevisions(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .reservations()
            .listRevisions(
                "276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
                "6ef59113-3482-40da-8d79-787f823e34bc",
                com.azure.core.util.Context.NONE);
    }
}
