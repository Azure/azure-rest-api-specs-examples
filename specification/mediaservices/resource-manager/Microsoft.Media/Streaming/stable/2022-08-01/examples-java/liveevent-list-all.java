/** Samples for LiveEvents List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-list-all.json
     */
    /**
     * Sample code: List all LiveEvents.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllLiveEvents(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.liveEvents().list("mediaresources", "slitestmedia10", com.azure.core.util.Context.NONE);
    }
}
