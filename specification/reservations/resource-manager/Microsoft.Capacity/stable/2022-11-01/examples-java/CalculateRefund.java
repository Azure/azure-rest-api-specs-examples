
import com.azure.resourcemanager.reservations.models.CalculateRefundRequest;
import com.azure.resourcemanager.reservations.models.CalculateRefundRequestProperties;
import com.azure.resourcemanager.reservations.models.ReservationToReturn;

/**
 * Samples for CalculateRefund Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/CalculateRefund.json
     */
    /**
     * Sample code: CalculateRefund.
     * 
     * @param manager Entry point to ReservationsManager.
     */
    public static void calculateRefund(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.calculateRefunds().postWithResponse("276e7ae4-84d0-4da6-ab4b-d6b94f3557da", new CalculateRefundRequest()
            .withId("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004")
            .withProperties(new CalculateRefundRequestProperties().withScope("Reservation")
                .withReservationToReturn(new ReservationToReturn().withReservationId(
                    "/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000")
                    .withQuantity(1))),
            com.azure.core.util.Context.NONE);
    }
}
