```java
import com.azure.core.util.Context;

/** Samples for Marketplaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesByDepartmentList.json
     */
    /**
     * Sample code: DepartmentMarketplacesList.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void departmentMarketplacesList(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.marketplaces().list("providers/Microsoft.Billing/departments/123456", null, null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
