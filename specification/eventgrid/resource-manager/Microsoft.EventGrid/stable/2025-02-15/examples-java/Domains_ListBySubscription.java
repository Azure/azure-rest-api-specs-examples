
/**
 * Samples for Domains List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * Domains_ListBySubscription.json
     */
    /**
     * Sample code: Domains_ListBySubscription.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().list(null, null, com.azure.core.util.Context.NONE);
    }
}
