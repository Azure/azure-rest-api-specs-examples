
/**
 * Samples for Reservation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationDetails.
     * json
     */
    /**
     * Sample code: GetReservation.
     * 
     * @param manager Entry point to ReservationsManager.
     */
    public static void getReservation(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.reservations().getWithResponse("276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
            "6ef59113-3482-40da-8d79-787f823e34bc", "renewProperties", com.azure.core.util.Context.NONE);
    }
}
