/** Samples for NetworkInterfaces ListEffectiveNetworkSecurityGroups. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/NetworkInterfaceEffectiveNSGList.json
     */
    /**
     * Sample code: List network interface effective network security groups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkInterfaceEffectiveNetworkSecurityGroups(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listEffectiveNetworkSecurityGroups("rg1", "nic1", com.azure.core.util.Context.NONE);
    }
}
