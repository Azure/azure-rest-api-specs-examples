
/**
 * Samples for ReservationRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationRecommendationsByResourceGroup.json
     */
    /**
     * Sample code: ReservationRecommendationsByResourceGroup-Legacy.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsByResourceGroupLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationRecommendations().list(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testGroup", null,
            com.azure.core.util.Context.NONE);
    }
}
