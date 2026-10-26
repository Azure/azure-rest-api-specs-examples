
/**
 * Samples for FlowProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getFlowProfile.json
     */
    /**
     * Sample code: Gets flowprofile resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void getsFlowprofileResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flowProfiles().getWithResponse("testRG", "testPipeline", "testFlowProfile",
            com.azure.core.util.Context.NONE);
    }
}
