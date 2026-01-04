
/**
 * Samples for NetworkSecurityPerimeterAssociations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspAssociationDelete.json
     */
    /**
     * Sample code: NspAssociationDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAssociationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAssociations().delete("rg1", "nsp1",
            "association1", com.azure.core.util.Context.NONE);
    }
}
