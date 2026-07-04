
/**
 * Samples for NetworkInterfaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceList.json
     */
    /**
     * Sample code: List network interfaces in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listNetworkInterfacesInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
