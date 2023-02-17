/** Samples for LiveEvents ListGetStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/liveevent-get-status.json
     */
    /**
     * Sample code: Get status of a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveEvents()
            .listGetStatus("mediaresources", "slitestmedia10", "myLiveEvent1", com.azure.core.util.Context.NONE);
    }
}
