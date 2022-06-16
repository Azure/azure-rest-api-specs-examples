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
