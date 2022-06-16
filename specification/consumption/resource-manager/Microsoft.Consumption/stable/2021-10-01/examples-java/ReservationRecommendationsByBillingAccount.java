import com.azure.core.util.Context;

/** Samples for ReservationRecommendations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsByBillingAccount.json
     */
    /**
     * Sample code: ReservationRecommendationsByBillingAccount-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsByBillingAccountLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationRecommendations()
            .list("providers/Microsoft.Billing/billingAccounts/123456", null, Context.NONE);
    }
}
