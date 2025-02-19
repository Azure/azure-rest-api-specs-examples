
/**
 * Samples for Reservation List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/
     * GetReservationsFromOrder.json
     */
    /**
     * Sample code: ReservationList.
     * 
     * @param manager Entry point to ReservationsManager.
     */
    public static void reservationList(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.reservations().list("276e7ae4-84d0-4da6-ab4b-d6b94f3557da", com.azure.core.util.Context.NONE);
    }
}
