import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.LookBackPeriod;
import com.azure.resourcemanager.consumption.models.Term;

/** Samples for ReservationRecommendationDetails Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationDetailsByBillingAccount.json
     */
    /**
     * Sample code: ReservationRecommendationsByBillingAccount-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsByBillingAccountLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationRecommendationDetails()
            .getWithResponse("Shared", "eastus", Term.P1Y, LookBackPeriod.LAST60DAYS, "Standard_DS14_v2", Context.NONE);
    }
}
