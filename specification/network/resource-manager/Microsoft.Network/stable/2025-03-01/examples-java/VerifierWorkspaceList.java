
/**
 * Samples for VerifierWorkspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VerifierWorkspaceList.json
     */
    /**
     * Sample code: VerifierWorkspaceList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void verifierWorkspaceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVerifierWorkspaces().list("rg1", "testNetworkManager", null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
