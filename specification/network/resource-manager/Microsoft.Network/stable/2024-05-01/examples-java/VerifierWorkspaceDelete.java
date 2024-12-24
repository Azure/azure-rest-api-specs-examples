
/**
 * Samples for VerifierWorkspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VerifierWorkspaceDelete.json
     */
    /**
     * Sample code: VerifierWorkspaceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void verifierWorkspaceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVerifierWorkspaces().delete("rg1", "testNetworkManager",
            "testWorkspace", com.azure.core.util.Context.NONE);
    }
}
