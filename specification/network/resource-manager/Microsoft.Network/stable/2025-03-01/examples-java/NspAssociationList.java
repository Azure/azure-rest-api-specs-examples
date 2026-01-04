
/**
 * Samples for NetworkSecurityPerimeterAssociations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspAssociationList.json
     */
    /**
     * Sample code: NspAssociationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAssociationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAssociations().list("rg1", "nsp1", null,
            null, com.azure.core.util.Context.NONE);
    }
}
