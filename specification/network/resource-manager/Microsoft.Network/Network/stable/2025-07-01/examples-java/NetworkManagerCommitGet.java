
/**
 * Samples for Commits Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerCommitGet.json
     */
    /**
     * Sample code: Get Network Manager Commit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNetworkManagerCommit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCommits().getWithResponse("myResourceGroup", "testNetworkManager", "myTestCommit",
            com.azure.core.util.Context.NONE);
    }
}
