
/**
 * Samples for VirtualNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualNetworks_List.json
     */
    /**
     * Sample code: VirtualNetworks_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualNetworksList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualNetworks().list("resourceGroupName", "{labName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
