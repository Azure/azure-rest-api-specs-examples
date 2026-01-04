
/**
 * Samples for NetworkSecurityPerimeterLinkReferences Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspLinkReferenceDelete.json
     */
    /**
     * Sample code: NspLinkReferenceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLinkReferenceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLinkReferences().delete("rg1", "nsp2",
            "link1-guid", com.azure.core.util.Context.NONE);
    }
}
