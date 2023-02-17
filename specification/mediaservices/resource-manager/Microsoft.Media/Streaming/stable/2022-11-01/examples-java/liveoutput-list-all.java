/** Samples for LiveOutputs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/liveoutput-list-all.json
     */
    /**
     * Sample code: List all LiveOutputs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllLiveOutputs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveOutputs()
            .list("mediaresources", "slitestmedia10", "myLiveEvent1", com.azure.core.util.Context.NONE);
    }
}
