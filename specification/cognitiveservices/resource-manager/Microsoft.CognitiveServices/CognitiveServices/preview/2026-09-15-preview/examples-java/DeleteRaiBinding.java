
/**
 * Samples for RaiBindings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/DeleteRaiBinding.json
     */
    /**
     * Sample code: Delete a RAI binding.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteARAIBinding(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBindings().deleteWithResponse("resource-group", "safety-account", "chat-guard",
            "\"00000000-0000-0000-0000-000000000002\"", com.azure.core.util.Context.NONE);
    }
}
