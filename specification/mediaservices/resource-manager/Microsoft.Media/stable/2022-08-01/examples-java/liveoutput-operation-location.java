import com.azure.core.util.Context;

/** Samples for LiveOutputs OperationLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveoutput-operation-location.json
     */
    /**
     * Sample code: Get the LiveOutput operation status.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getTheLiveOutputOperationStatus(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveOutputs()
            .operationLocationWithResponse(
                "mediaresources",
                "slitestmedia10",
                "myLiveEvent1",
                "myLiveOutput1",
                "62e4d893-d233-4005-988e-a428d9f77076",
                Context.NONE);
    }
}
