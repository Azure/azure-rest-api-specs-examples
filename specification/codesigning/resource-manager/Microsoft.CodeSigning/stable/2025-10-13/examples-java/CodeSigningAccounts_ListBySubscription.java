
/**
 * Samples for CodeSigningAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_ListBySubscription.json
     */
    /**
     * Sample code: Lists artifact signing accounts within a subscription.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void listsArtifactSigningAccountsWithinASubscription(
        com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.codeSigningAccounts().list(com.azure.core.util.Context.NONE);
    }
}
