import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.LiveEventActionInput;

/** Samples for LiveEvents Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-stop.json
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
                Context.NONE);
    }
}
