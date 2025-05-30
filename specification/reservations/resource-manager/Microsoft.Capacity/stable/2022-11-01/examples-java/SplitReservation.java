
import com.azure.resourcemanager.reservations.models.SplitRequest;
import java.util.Arrays;

/**
 * Samples for Reservation Split.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/SplitReservation.json
     */
    /**
     * Sample code: Split.
     * 
     * @param manager Entry point to ReservationsManager.
     */
    public static void split(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.reservations().split("276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
            new SplitRequest().withQuantities(Arrays.asList(1, 2)).withReservationId(
                "/providers/Microsoft.Capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/bcae77cd-3119-4766-919f-b50d36c75c7a"),
            com.azure.core.util.Context.NONE);
    }
}
