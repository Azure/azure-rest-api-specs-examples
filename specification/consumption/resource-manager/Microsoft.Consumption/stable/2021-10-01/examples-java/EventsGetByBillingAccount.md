```java
import com.azure.core.util.Context;

/** Samples for EventsOperation ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/EventsGetByBillingAccount.json
     */
    /**
     * Sample code: EventsGetByBillingAccount.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void eventsGetByBillingAccount(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.eventsOperations().listByBillingAccount("1234:5678", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.
