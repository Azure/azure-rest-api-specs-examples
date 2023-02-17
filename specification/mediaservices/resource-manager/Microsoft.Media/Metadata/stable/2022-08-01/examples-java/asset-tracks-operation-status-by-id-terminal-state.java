/** Samples for OperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-operation-status-by-id-terminal-state.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is completed.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsCompleted(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .operationStatuses()
            .getWithResponse(
                "contosorg",
                "contosomedia",
                "ClimbingMountRainer",
                "text1",
                "e78f8d40-7aaa-4f2f-8ae6-73987e7c5a08",
                com.azure.core.util.Context.NONE);
    }
}
