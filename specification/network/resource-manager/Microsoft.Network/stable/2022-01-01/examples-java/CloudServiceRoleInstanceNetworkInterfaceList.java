import com.azure.core.util.Context;

/** Samples for NetworkInterfaces ListCloudServiceRoleInstanceNetworkInterfaces. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/CloudServiceRoleInstanceNetworkInterfaceList.json
     */
    /**
     * Sample code: List cloud service role instance network interfaces.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServiceRoleInstanceNetworkInterfaces(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listCloudServiceRoleInstanceNetworkInterfaces("rg1", "cs1", "TestVMRole_IN_0", Context.NONE);
    }
}
