import com.azure.core.util.Context;

/** Samples for LiveEvents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-list-by-name.json
     */
    /**
     * Sample code: Get a LiveEvent by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getALiveEventByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.liveEvents().getWithResponse("mediaresources", "slitestmedia10", "myLiveEvent1", Context.NONE);
    }
}
