
/**
 * Samples for ReservationRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationRecommendationsFilterBySubscriptionForScopeLookBackPeriod.json
     */
    /**
     * Sample code: ReservationRecommendationsFilterBySubscriptionForScopeLookBackPeriod-Legacy.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsFilterBySubscriptionForScopeLookBackPeriodLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationRecommendations().list("subscriptions/00000000-0000-0000-0000-000000000000",
            "properties/scope eq 'Single' AND properties/lookBackPeriod eq 'Last7Days'",
            com.azure.core.util.Context.NONE);
    }
}
