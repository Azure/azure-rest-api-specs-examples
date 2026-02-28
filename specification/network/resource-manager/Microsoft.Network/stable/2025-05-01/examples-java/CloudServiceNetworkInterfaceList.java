
/**
 * Samples for NetworkInterfaces ListCloudServiceNetworkInterfaces.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * CloudServiceNetworkInterfaceList.json
     */
    /**
     * Sample code: List cloud service network interfaces.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServiceNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces().listCloudServiceNetworkInterfaces("rg1",
            "cs1", com.azure.core.util.Context.NONE);
    }
}
