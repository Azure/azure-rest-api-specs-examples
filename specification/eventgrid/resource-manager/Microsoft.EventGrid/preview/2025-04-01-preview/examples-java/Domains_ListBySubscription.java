
/**
 * Samples for Domains List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
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
