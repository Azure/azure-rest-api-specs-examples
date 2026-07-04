
/**
 * Samples for Commits List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerCommitList.json
     */
    /**
     * Sample code: List Network Manager Commit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listNetworkManagerCommit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCommits().list("myResourceGroup", "testNetworkManager", null, null,
            com.azure.core.util.Context.NONE);
    }
}
