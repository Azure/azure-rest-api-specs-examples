
/**
 * Samples for EventCategories List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetEventCategories.json
     */
    /**
     * Sample code: Get event categories.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getEventCategories(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getEventCategories()
            .list(com.azure.core.util.Context.NONE);
    }
}
