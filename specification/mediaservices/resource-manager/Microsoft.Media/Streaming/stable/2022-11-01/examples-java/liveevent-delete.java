/** Samples for LiveEvents Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/liveevent-delete.json
     */
    /**
     * Sample code: Delete a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveEvents()
            .delete("mediaresources", "slitestmedia10", "myLiveEvent1", com.azure.core.util.Context.NONE);
    }
}
