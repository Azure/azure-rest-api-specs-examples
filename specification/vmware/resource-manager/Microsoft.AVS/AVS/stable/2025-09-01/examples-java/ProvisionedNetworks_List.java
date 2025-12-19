
/**
 * Samples for ProvisionedNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ProvisionedNetworks_List.json
     */
    /**
     * Sample code: ProvisionedNetworks_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void provisionedNetworksList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.provisionedNetworks().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
