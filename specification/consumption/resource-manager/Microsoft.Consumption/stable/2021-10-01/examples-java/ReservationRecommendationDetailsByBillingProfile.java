
import com.azure.resourcemanager.consumption.models.LookBackPeriod;
import com.azure.resourcemanager.consumption.models.Scope;
import com.azure.resourcemanager.consumption.models.Term;

/**
 * Samples for ReservationRecommendationDetails Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationRecommendationDetailsByBillingProfile.json
     */
    /**
     * Sample code: ReservationRecommendationsByBillingProfile-Modern.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsByBillingProfileModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationRecommendationDetails().getWithResponse(
            "providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-00000000:00000000-0000-0000-0000-00000000/billingProfiles/00000000-0000-0000-0000-00000000",
            Scope.SHARED, "australiaeast", Term.P1Y, LookBackPeriod.LAST7DAYS, "Standard_B2s",
            com.azure.core.util.Context.NONE);
    }
}
