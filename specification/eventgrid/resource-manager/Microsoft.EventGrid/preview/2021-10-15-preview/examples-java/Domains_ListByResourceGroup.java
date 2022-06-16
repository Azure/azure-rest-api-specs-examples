import com.azure.core.util.Context;

/** Samples for Domains ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Domains_ListByResourceGroup.json
     */
    /**
     * Sample code: Domains_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
