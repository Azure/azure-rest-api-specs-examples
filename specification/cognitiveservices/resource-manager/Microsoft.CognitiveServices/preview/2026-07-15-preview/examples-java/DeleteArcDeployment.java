
/**
 * Samples for ArcDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15-preview/DeleteArcDeployment.json
     */
    /**
     * Sample code: DeleteArcDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteArcDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.arcDeployments().delete("resourceGroupName", "accountName", "phi-3-arc",
            com.azure.core.util.Context.NONE);
    }
}
