
/**
 * Samples for MediaServicesOperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/media-service-
     * operation-status-by-id-non-terminal-state.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is ongoing.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsOngoing(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.mediaServicesOperationStatuses().getWithResponse("westus", "D612C429-2526-49D5-961B-885AE11406FD",
            com.azure.core.util.Context.NONE);
    }
}
