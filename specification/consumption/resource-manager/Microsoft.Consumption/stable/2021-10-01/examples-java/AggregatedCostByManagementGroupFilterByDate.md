Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AggregatedCost GetByManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostByManagementGroupFilterByDate.json
     */
    /**
     * Sample code: AggregatedCostByManagementGroupFilterByDate.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void aggregatedCostByManagementGroupFilterByDate(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .aggregatedCosts()
            .getByManagementGroupWithResponse(
                "managementGroupForTest",
                "usageStart ge '2018-08-15' and properties/usageStart le '2018-08-31'",
                Context.NONE);
    }
}
```
