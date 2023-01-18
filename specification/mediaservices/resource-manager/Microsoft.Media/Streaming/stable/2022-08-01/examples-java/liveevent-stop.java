import com.azure.resourcemanager.mediaservices.models.LiveEventActionInput;

/** Samples for LiveEvents Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-stop.json
     */
    /**
     * Sample code: Stop a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void stopALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveEvents()
            .stop(
                "mediaresources",
                "slitestmedia10",
                "myLiveEvent1",
                new LiveEventActionInput().withRemoveOutputsOnStop(false),
                com.azure.core.util.Context.NONE);
    }
}
