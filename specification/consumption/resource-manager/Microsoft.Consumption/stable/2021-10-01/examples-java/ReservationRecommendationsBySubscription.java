import com.azure.core.util.Context;

/** Samples for ReservationRecommendations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsBySubscription.json
     */
    /**
     * Sample code: ReservationRecommendationsBySubscription-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsBySubscriptionLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationRecommendations()
            .list("subscriptions/00000000-0000-0000-0000-000000000000", null, Context.NONE);
    }
}
