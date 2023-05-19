/** Samples for Domains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Domains_Delete.json
     */
    /**
     * Sample code: Domains_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().delete("examplerg", "exampledomain1", com.azure.core.util.Context.NONE);
    }
}
