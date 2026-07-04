
/**
 * Samples for VerifierWorkspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VerifierWorkspacePatch.json
     */
    /**
     * Sample code: VerifierWorkspacePatch.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void verifierWorkspacePatch(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVerifierWorkspaces().updateWithResponse("rg1", "testNetworkManager", "testWorkspace",
            null, null, com.azure.core.util.Context.NONE);
    }
}
