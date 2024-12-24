
/**
 * Samples for VerifierWorkspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VerifierWorkspacePatch.json
     */
    /**
     * Sample code: VerifierWorkspacePatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void verifierWorkspacePatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVerifierWorkspaces().updateWithResponse("rg1",
            "testNetworkManager", "testWorkspace", null, com.azure.core.util.Context.NONE);
    }
}
