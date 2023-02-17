/** Samples for LiveEvents ListGetStreamEvents. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/liveevent-get-stream-events.json
     */
    /**
     * Sample code: Get stream events of a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStreamEventsOfALiveEvent(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveEvents()
            .listGetStreamEvents("mediaresources", "slitestmedia10", "myLiveEvent1", com.azure.core.util.Context.NONE);
    }
}
