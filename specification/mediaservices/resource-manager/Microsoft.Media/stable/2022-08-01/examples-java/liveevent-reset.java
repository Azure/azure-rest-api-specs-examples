import com.azure.core.util.Context;

/** Samples for LiveEvents Reset. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveevent-reset.json
     */
    /**
     * Sample code: Reset a LiveEvent.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void resetALiveEvent(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.liveEvents().reset("mediaresources", "slitestmedia10", "myLiveEvent1", Context.NONE);
    }
}
