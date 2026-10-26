
/**
 * Samples for FlowProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/deleteFlowProfile.json
     */
    /**
     * Sample code: Deletes the flowprofile resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        deletesTheFlowprofileResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flowProfiles().delete("testRG", "testPipeline", "testFlowProfile", com.azure.core.util.Context.NONE);
    }
}
