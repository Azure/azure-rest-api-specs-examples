Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Datagrain;

/** Samples for ReservationsSummaries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthlyWithBillingProfileIdReservationId.json
     */
    /**
     * Sample code: ReservationSummariesMonthlyWithBillingProfileIdReservationId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationSummariesMonthlyWithBillingProfileIdReservationId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsSummaries()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579",
                Datagrain.MONTHLY,
                null,
                null,
                null,
                "1c6b6358-709f-484c-85f1-72e862a0cf3b",
                "9f39ba10-794f-4dcb-8f4b-8d0cb47c27dc",
                Context.NONE);
    }
}
```
