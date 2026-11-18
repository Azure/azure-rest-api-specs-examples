
import com.azure.resourcemanager.cognitiveservices.models.DeleteRequestClusterDisposition;

/**
 * Samples for AttachedAksClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/DeleteAttachedAksCluster.json
     */
    /**
     * Sample code: DeleteAttachedAksCluster.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteAttachedAksCluster(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.attachedAksClusters().deleteWithResponse("rgcognitiveservices", "myAccount", "myAttachment",
            DeleteRequestClusterDisposition.KEEP, com.azure.core.util.Context.NONE);
    }
}
