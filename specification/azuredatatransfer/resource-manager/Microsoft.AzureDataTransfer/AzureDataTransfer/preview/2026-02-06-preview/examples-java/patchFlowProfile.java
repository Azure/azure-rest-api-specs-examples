
import com.azure.resourcemanager.datatransfer.models.FlowProfile;

/**
 * Samples for FlowProfiles Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/patchFlowProfile.json
     */
    /**
     * Sample code: Updates the flowprofile resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        updatesTheFlowprofileResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        FlowProfile resource = manager.flowProfiles()
            .getWithResponse("testRG", "testPipeline", "testFlowProfile", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
