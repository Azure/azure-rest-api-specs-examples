
/**
 * Samples for RaiBindings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/GetRaiBinding.json
     */
    /**
     * Sample code: Get a RAI binding.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getARAIBinding(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBindings().getWithResponse("resource-group", "safety-account", "chat-guard",
            com.azure.core.util.Context.NONE);
    }
}
