import com.azure.resourcemanager.reservations.models.RefundRequest;
import com.azure.resourcemanager.reservations.models.RefundRequestProperties;
import com.azure.resourcemanager.reservations.models.ReservationToReturn;

/** Samples for Return Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/Return.json
     */
    /**
     * Sample code: Return a reservation.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void returnAReservation(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .returns()
            .post(
                "50000000-aaaa-bbbb-cccc-100000000004",
                new RefundRequest()
                    .withProperties(
                        new RefundRequestProperties()
                            .withSessionId("10000000-aaaa-bbbb-cccc-200000000000")
                            .withScope("Reservation")
                            .withReservationToReturn(
                                new ReservationToReturn()
                                    .withReservationId(
                                        "/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000")
                                    .withQuantity(1))
                            .withReturnReason("PurchasedWrongProduct")),
                com.azure.core.util.Context.NONE);
    }
}
