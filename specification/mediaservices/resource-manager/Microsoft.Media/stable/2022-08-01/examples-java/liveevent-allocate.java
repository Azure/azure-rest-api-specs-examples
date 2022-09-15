import com.azure.core.util.Context;

/** Samples for LiveEvents Allocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveevent-allocate.json
     */
    /**
     * Sample code: Allocate a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void allocateALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.liveEvents().allocate("mediaresources", "slitestmedia10", "myLiveEvent1", Context.NONE);
    }
}
