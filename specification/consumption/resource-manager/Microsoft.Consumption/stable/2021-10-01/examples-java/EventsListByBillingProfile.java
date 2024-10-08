
/**
 * Samples for EventsOperation ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * EventsListByBillingProfile.json
     */
    /**
     * Sample code: EventsListByBillingProfile.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void eventsListByBillingProfile(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.eventsOperations().listByBillingProfile("1234:5678", "4268", "2019-09-01", "2019-10-31",
            com.azure.core.util.Context.NONE);
    }
}
