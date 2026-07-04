
/**
 * Samples for VerifierWorkspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VerifierWorkspaceDelete.json
     */
    /**
     * Sample code: VerifierWorkspaceDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void verifierWorkspaceDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVerifierWorkspaces().delete("rg1", "testNetworkManager", "testWorkspace", null,
            com.azure.core.util.Context.NONE);
    }
}
