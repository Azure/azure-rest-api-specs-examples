/** Samples for EventsOperation ListByTenantId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Events_ListByTenantId.json
     */
    /**
     * Sample code: ListEventsByTenantId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listEventsByTenantId(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .eventsOperations()
            .listByTenantId(
                "service eq 'Virtual Machines' or region eq 'West US'", "7/24/2020", com.azure.core.util.Context.NONE);
    }
}
