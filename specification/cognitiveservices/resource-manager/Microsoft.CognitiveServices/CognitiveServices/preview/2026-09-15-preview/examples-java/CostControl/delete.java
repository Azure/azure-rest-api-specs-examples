
/**
 * Samples for CostControls Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CostControl/delete.json
     */
    /**
     * Sample code: Delete a cost control conditionally.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteACostControlConditionally(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.costControls().deleteWithResponse("foundry-resource-group", "foundry-account", "production-agents",
            "\"00000000-0000-0000-0000-000000000002\"", com.azure.core.util.Context.NONE);
    }
}
