Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Budgets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BudgetsList.json
     */
    /**
     * Sample code: BudgetsList.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void budgetsList(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.budgets().list("subscriptions/00000000-0000-0000-0000-000000000000", Context.NONE);
    }
}
```
