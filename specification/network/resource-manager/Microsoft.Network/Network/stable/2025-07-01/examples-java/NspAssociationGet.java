
/**
 * Samples for NetworkSecurityPerimeterAssociations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAssociationGet.json
     */
    /**
     * Sample code: NspAssociationGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAssociationGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAssociations().getWithResponse("rg1", "nsp1", "association1",
            com.azure.core.util.Context.NONE);
    }
}
