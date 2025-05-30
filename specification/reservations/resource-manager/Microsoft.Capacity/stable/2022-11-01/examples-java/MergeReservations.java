
import com.azure.resourcemanager.reservations.models.MergeRequest;
import java.util.Arrays;

/**
 * Samples for Reservation Merge.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/MergeReservations.json
     */
    /**
     * Sample code: Merge.
     * 
     * @param manager Entry point to ReservationsManager.
     */
    public static void merge(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.reservations().merge("276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
            new MergeRequest().withSources(Arrays.asList(
                "/providers/Microsoft.Capacity/reservationOrders/c0565a8a-4491-4e77-b07b-5e6d66718e1c/reservations/cea04232-932e-47db-acb5-e29a945ecc73",
                "/providers/Microsoft.Capacity/reservationOrders/c0565a8a-4491-4e77-b07b-5e6d66718e1c/reservations/5bf54dc7-dacd-4f46-a16b-7b78f4a59799")),
            com.azure.core.util.Context.NONE);
    }
}
