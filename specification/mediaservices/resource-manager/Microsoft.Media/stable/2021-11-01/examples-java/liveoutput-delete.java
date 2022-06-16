import com.azure.core.util.Context;

/** Samples for LiveOutputs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveoutput-delete.json
     */
    /**
     * Sample code: Delete a LiveOutput.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteALiveOutput(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.liveOutputs().delete("mediaresources", "slitestmedia10", "myLiveEvent1", "myLiveOutput1", Context.NONE);
    }
}
