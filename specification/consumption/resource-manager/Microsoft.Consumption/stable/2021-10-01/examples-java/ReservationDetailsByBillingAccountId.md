```java
import com.azure.core.util.Context;

/** Samples for ReservationsDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetailsByBillingAccountId.json
     */
    /**
     * Sample code: ReservationDetailsByBillingAccountId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationDetailsByBillingAccountId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsDetails()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345",
                null,
                null,
                "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05",
                null,
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
