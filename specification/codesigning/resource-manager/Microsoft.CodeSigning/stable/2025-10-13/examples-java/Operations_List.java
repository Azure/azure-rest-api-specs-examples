
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/Operations_List.json
     */
    /**
     * Sample code: List artifact signing account operations.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void
        listArtifactSigningAccountOperations(com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
