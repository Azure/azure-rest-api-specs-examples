import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDailyWithBillingAccountId.json
     */
    /**
     * Sample code: ReservationSummariesDailyWithBillingAccountId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesDailyWithBillingAccountId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345",
                Datagrain.DAILY,
                null,
                null,
                "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-11-20",
                null,
                null,
                Context.NONE);
    }
}
