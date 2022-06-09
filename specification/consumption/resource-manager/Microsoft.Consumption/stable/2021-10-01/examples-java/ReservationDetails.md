```java
import com.azure.core.util.Context;

/** Samples for ReservationsDetails ListByReservationOrder. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetails.json
     */
    /**
     * Sample code: ReservationDetails.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationDetails(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsDetails()
            .listByReservationOrder(
                "00000000-0000-0000-0000-000000000000",
                "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
