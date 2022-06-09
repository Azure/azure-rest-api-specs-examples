```java
import com.azure.core.util.Context;

/** Samples for Tags Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Tags.json
     */
    /**
     * Sample code: Tags_Get.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void tagsGet(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.tags().getWithResponse("providers/Microsoft.CostManagement/billingAccounts/1234", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
