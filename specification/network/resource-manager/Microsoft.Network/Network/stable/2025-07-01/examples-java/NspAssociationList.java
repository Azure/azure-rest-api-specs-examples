
/**
 * Samples for NetworkSecurityPerimeterAssociations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAssociationList.json
     */
    /**
     * Sample code: NspAssociationList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAssociationList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAssociations().list("rg1", "nsp1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
