import com.azure.core.util.Context;

/** Samples for ReservationRecommendations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsByBillingProfile.json
     */
    /**
     * Sample code: ReservationRecommendationsByBillingProfile-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsByBillingProfileModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationRecommendations()
            .list("providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420", null, Context.NONE);
    }
}
