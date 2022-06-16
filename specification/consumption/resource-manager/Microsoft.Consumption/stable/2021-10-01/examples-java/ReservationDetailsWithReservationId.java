import com.azure.core.util.Context;

/** Samples for ReservationsDetails ListByReservationOrderAndReservation. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetailsWithReservationId.json
     */
    /**
     * Sample code: ReservationDetailsWithReservationId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationDetailsWithReservationId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsDetails()
            .listByReservationOrderAndReservation(
                "00000000-0000-0000-0000-000000000000",
                "00000000-0000-0000-0000-000000000000",
                "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05",
                Context.NONE);
    }
}
