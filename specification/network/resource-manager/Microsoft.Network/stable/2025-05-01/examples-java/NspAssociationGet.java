
/**
 * Samples for NetworkSecurityPerimeterAssociations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspAssociationGet.json
     */
    /**
     * Sample code: NspAssociationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAssociationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAssociations().getWithResponse("rg1",
            "nsp1", "association1", com.azure.core.util.Context.NONE);
    }
}
