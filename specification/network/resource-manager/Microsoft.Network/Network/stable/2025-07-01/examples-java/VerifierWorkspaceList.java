
/**
 * Samples for VerifierWorkspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VerifierWorkspaceList.json
     */
    /**
     * Sample code: VerifierWorkspaceList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void verifierWorkspaceList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVerifierWorkspaces().list("rg1", "testNetworkManager", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
