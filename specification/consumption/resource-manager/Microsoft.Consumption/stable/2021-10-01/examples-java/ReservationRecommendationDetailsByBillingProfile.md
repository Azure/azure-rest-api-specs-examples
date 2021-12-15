Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.LookBackPeriod;
import com.azure.resourcemanager.consumption.models.Term;

/** Samples for ReservationRecommendationDetails Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationDetailsByBillingProfile.json
     */
    /**
     * Sample code: ReservationRecommendationsByBillingProfile-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationRecommendationsByBillingProfileModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationRecommendationDetails()
            .getWithResponse(
                "Shared", "australiaeast", Term.P1Y, LookBackPeriod.LAST7DAYS, "Standard_B2s", Context.NONE);
    }
}
```
