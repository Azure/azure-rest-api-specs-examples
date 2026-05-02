
/**
 * Samples for RaiToolLabels Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetRaiToolLabel.json
     */
    /**
     * Sample code: GetRaiToolLabel.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getRaiToolLabel(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiToolLabels().getWithResponse("resourceGroupName", "accountName", "ToolLabelName",
            com.azure.core.util.Context.NONE);
    }
}
