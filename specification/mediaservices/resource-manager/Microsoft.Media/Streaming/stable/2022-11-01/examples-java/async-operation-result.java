
/**
 * Samples for LiveOutputs AsyncOperation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/async-operation
     * -result.json
     */
    /**
     * Sample code: Get the LiveOutput operation status.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void
        getTheLiveOutputOperationStatus(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.liveOutputs().asyncOperationWithResponse("mediaresources", "slitestmedia10",
            "62e4d893-d233-4005-988e-a428d9f77076", com.azure.core.util.Context.NONE);
    }
}
