
/**
 * Samples for CodeSigningAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_Get.json
     */
    /**
     * Sample code: Get an artifact signing account.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void
        getAnArtifactSigningAccount(com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.codeSigningAccounts().getByResourceGroupWithResponse("MyResourceGroup", "MyAccount",
            com.azure.core.util.Context.NONE);
    }
}
