
/**
 * Samples for OperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-
     * operation-status-by-id-terminal-state-failed.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is completed with error.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsCompletedWithError(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.operationStatuses().getWithResponse("contosorg", "contosomedia", "ClimbingMountRainer", "text1",
            "86835197-3b47-402e-b313-70b82eaba296", com.azure.core.util.Context.NONE);
    }
}
