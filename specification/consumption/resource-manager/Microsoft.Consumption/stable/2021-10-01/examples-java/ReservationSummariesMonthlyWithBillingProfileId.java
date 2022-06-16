import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthlyWithBillingProfileId.json
     */
    /**
     * Sample code: ReservationSummariesMonthlyWithBillingProfileId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesMonthlyWithBillingProfileId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579",
                Datagrain.MONTHLY,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
