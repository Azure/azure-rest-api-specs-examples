
/**
 * Samples for NetworkSecurityPerimeterAssociations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAssociationDelete.json
     */
    /**
     * Sample code: NspAssociationDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAssociationDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAssociations().delete("rg1", "nsp1", "association1",
            com.azure.core.util.Context.NONE);
    }
}
