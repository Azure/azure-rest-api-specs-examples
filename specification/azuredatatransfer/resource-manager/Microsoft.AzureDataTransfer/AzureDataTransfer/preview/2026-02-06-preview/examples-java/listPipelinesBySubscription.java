
/**
 * Samples for Pipelines ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listPipelinesBySubscription.json
     */
    /**
     * Sample code: Gets pipelines in a subscription.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getsPipelinesInASubscription(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().listBySubscription(com.azure.core.util.Context.NONE);
    }
}
