
/**
 * Samples for ArcDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/DeleteArcDeploymentWithForce.json
     */
    /**
     * Sample code: DeleteArcDeploymentWithForce.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteArcDeploymentWithForce(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.arcDeployments().delete("resourceGroupName", "accountName", "phi-3-arc", true,
            com.azure.core.util.Context.NONE);
    }
}
