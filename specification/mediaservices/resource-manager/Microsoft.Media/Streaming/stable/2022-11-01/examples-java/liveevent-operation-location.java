/** Samples for LiveEvents OperationLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/liveevent-operation-location.json
     */
    /**
     * Sample code: Get the LiveEvent operation status.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getTheLiveEventOperationStatus(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .liveEvents()
            .operationLocationWithResponse(
                "mediaresources",
                "slitestmedia10",
                "myLiveEvent1",
                "62e4d893-d233-4005-988e-a428d9f77076",
                com.azure.core.util.Context.NONE);
    }
}
