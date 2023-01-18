/** Samples for OperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-operation-status-by-id-non-terminal-state.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is ongoing.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsOngoing(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .operationStatuses()
            .getWithResponse(
                "contoso",
                "contosomedia",
                "ClimbingMountRainer",
                "text1",
                "5827d9a1-1fb4-4e54-ac40-8eeed9b862c8",
                com.azure.core.util.Context.NONE);
    }
}
