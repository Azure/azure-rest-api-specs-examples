
import com.azure.resourcemanager.consumption.models.Datagrain;

/**
 * Samples for ReservationsSummaries ListByReservationOrderAndReservation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationSummariesDailyWithReservationId.json
     */
    /**
     * Sample code: ReservationSummariesDailyWithReservationId.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void
        reservationSummariesDailyWithReservationId(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationsSummaries().listByReservationOrderAndReservation("00000000-0000-0000-0000-000000000000",
            "00000000-0000-0000-0000-000000000000", Datagrain.DAILY,
            "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-11-20",
            com.azure.core.util.Context.NONE);
    }
}
