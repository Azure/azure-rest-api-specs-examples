import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries ListByReservationOrder. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDaily.json
     */
    /**
     * Sample code: ReservationSummariesDaily.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesDaily(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .listByReservationOrder(
                "00000000-0000-0000-0000-000000000000",
                Datagrain.DAILY,
                "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-11-20",
                Context.NONE);
    }
}
