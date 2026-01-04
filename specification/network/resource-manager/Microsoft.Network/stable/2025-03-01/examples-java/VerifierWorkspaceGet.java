
/**
 * Samples for VerifierWorkspaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VerifierWorkspaceGet.json
     */
    /**
     * Sample code: VerifierWorkspaceGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void verifierWorkspaceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVerifierWorkspaces().getWithResponse("rg1", "testNetworkManager",
            "testWorkspace", com.azure.core.util.Context.NONE);
    }
}
