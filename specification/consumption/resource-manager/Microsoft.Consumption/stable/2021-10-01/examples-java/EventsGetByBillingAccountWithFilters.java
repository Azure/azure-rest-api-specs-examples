import com.azure.core.util.Context;

/** Samples for EventsOperation ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/EventsGetByBillingAccountWithFilters.json
     */
    /**
     * Sample code: EventsGetByBillingAccountWithFilters.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void eventsGetByBillingAccountWithFilters(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .eventsOperations()
            .listByBillingAccount(
                "1234:5678", "lotid eq 'G202001083926600XXXXX' AND lotsource eq 'consumptioncommitment'", Context.NONE);
    }
}
