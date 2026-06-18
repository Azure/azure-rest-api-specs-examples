
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/Operations_List.json
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
