
/**
 * Samples for NetworkSecurityPerimeterLinkReferences List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspLinkReferenceList.json
     */
    /**
     * Sample code: NspLinkReferenceList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLinkReferenceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLinkReferences().list("rg1", "nsp2", null,
            null, com.azure.core.util.Context.NONE);
    }
}
