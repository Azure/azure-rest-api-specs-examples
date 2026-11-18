
/**
 * Samples for AttachedAksClusters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/GetAttachedAksCluster.json
     */
    /**
     * Sample code: GetAttachedAksCluster.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAttachedAksCluster(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.attachedAksClusters().getWithResponse("rgcognitiveservices", "myAccount", "myAttachment",
            com.azure.core.util.Context.NONE);
    }
}
