```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Metrictype;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListByMetricActualCost.json
     */
    /**
     * Sample code: UsageDetailsListByMetricActualCost-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void usageDetailsListByMetricActualCostLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                null,
                null,
                null,
                null,
                Metrictype.ACTUALCOST,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
