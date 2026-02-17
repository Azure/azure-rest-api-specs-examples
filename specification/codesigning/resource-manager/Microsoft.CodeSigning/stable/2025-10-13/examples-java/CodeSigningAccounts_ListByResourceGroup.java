
/**
 * Samples for CodeSigningAccounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists artifact signing accounts within a resource group.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void listsArtifactSigningAccountsWithinAResourceGroup(
        com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.codeSigningAccounts().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
