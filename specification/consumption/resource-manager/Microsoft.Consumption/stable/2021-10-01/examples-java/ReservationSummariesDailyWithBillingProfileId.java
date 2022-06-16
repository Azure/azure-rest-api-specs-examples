import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDailyWithBillingProfileId.json
     */
    /**
     * Sample code: ReservationSummariesDailyWithBillingProfileId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesDailyWithBillingProfileId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579",
                Datagrain.DAILY,
                "2017-10-01",
                "2017-11-20",
                null,
                null,
                null,
                Context.NONE);
    }
}
