```java
import com.azure.core.util.Context;

/** Samples for Budgets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
     */
    /**
     * Sample code: DeleteBudget.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void deleteBudget(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .budgets()
            .deleteWithResponse("subscriptions/00000000-0000-0000-0000-000000000000", "TestBudget", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
