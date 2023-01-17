/** Samples for LiveOutputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-list-by-name.json
     */
    /**
     * Sample code: Get a LiveOutput by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getALiveOutputByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveOutputs()
            .getWithResponse(
                "mediaresources", "slitestmedia10", "myLiveEvent1", "myLiveOutput1", com.azure.core.util.Context.NONE);
    }
}
