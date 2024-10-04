
import com.azure.resourcemanager.consumption.models.LookBackPeriod;
import com.azure.resourcemanager.consumption.models.Scope;
import com.azure.resourcemanager.consumption.models.Term;

/**
 * Samples for ReservationRecommendationDetails Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationRecommendationDetailsBySubscription.json
     */
    /**
     * Sample code: ReservationRecommendationsBySubscription-Legacy.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsBySubscriptionLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationRecommendationDetails().getWithResponse("subscriptions/00000000-0000-0000-0000-00000000",
            Scope.SINGLE, "westus", Term.P3Y, LookBackPeriod.LAST30DAYS, "Standard_DS13_v2",
            com.azure.core.util.Context.NONE);
    }
}
