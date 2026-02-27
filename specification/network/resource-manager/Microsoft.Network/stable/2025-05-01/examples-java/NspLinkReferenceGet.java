
/**
 * Samples for NetworkSecurityPerimeterLinkReferences Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLinkReferenceGet.json
     */
    /**
     * Sample code: NspLinkReferencesGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLinkReferencesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLinkReferences().getWithResponse("rg1",
            "nsp2", "link1-guid", com.azure.core.util.Context.NONE);
    }
}
