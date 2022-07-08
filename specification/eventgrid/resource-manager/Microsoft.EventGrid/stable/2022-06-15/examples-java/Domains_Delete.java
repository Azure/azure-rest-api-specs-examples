import com.azure.core.util.Context;

/** Samples for Domains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Domains_Delete.json
     */
    /**
     * Sample code: Domains_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().delete("examplerg", "exampledomain1", Context.NONE);
    }
}
