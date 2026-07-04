
/**
 * Samples for Commits Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerCommitDelete.json
     */
    /**
     * Sample code: Delete Network Manager Commit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkManagerCommit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCommits().delete("myResourceGroup", "testNetworkManager", "myTestCommit",
            com.azure.core.util.Context.NONE);
    }
}
