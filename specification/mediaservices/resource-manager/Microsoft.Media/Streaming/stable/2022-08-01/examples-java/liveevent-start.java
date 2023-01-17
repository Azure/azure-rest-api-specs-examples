/** Samples for LiveEvents Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-start.json
     */
    /**
     * Sample code: Start a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void startALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveEvents()
            .start("mediaresources", "slitestmedia10", "myLiveEvent1", com.azure.core.util.Context.NONE);
    }
}
