
/**
 * Samples for NetworkProfiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkProfileGetWithContainerNic.json
     */
    /**
     * Sample code: Get network profile with container network interfaces.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getNetworkProfileWithContainerNetworkInterfaces(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkProfiles().getByResourceGroupWithResponse("rg1", "networkProfile1", null,
            com.azure.core.util.Context.NONE);
    }
}
