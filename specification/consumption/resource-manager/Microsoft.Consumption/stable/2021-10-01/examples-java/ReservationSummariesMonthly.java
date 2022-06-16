import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries ListByReservationOrder. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthly.json
     */
    /**
     * Sample code: ReservationSummariesMonthly.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesMonthly(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .listByReservationOrder("00000000-0000-0000-0000-000000000000", Datagrain.MONTHLY, null, Context.NONE);
    }
}
