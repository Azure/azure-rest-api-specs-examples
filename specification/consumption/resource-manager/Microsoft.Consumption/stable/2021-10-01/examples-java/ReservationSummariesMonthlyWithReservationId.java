import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries ListByReservationOrderAndReservation. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthlyWithReservationId.json
     */
    /**
     * Sample code: ReservationSummariesMonthlyWithReservationId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesMonthlyWithReservationId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .listByReservationOrderAndReservation(
                "00000000-0000-0000-0000-000000000000",
                "00000000-0000-0000-0000-000000000000",
                Datagrain.MONTHLY,
                null,
                Context.NONE);
    }
}
