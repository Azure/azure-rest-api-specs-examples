/** Samples for Domains ListSharedAccessKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Domains_ListSharedAccessKeys.json
     */
    /**
     * Sample code: Domains_ListSharedAccessKeys.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsListSharedAccessKeys(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domains()
            .listSharedAccessKeysWithResponse("examplerg", "exampledomain2", com.azure.core.util.Context.NONE);
    }
}
