
/**
 * Samples for Channels GetFullUrl.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * Channels_GetFullUrl.json
     */
    /**
     * Sample code: Channels_GetFullUrl.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsGetFullUrl(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.channels().getFullUrlWithResponse("examplerg", "examplenamespace", "examplechannel",
            com.azure.core.util.Context.NONE);
    }
}
