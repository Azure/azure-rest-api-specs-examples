
/**
 * Samples for CodeSigningAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_Delete.json
     */
    /**
     * Sample code: Delete an artifact signing account.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void
        deleteAnArtifactSigningAccount(com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.codeSigningAccounts().delete("MyResourceGroup", "MyAccount", com.azure.core.util.Context.NONE);
    }
}
