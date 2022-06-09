```java
import com.azure.core.util.Context;

/** Samples for ReservationTransactions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationTransactionsListByEnrollmentNumber.json
     */
    /**
     * Sample code: ReservationTransactionsByEnrollmentNumber.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationTransactionsByEnrollmentNumber(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationTransactions()
            .list("123456", "properties/eventDate ge 2020-05-20 AND properties/eventDate le 2020-05-30", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
