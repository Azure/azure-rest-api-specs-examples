import com.azure.core.util.Context;

/** Samples for NetworkInterfaces GetCloudServiceNetworkInterface. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/CloudServiceNetworkInterfaceGet.json
     */
    /**
     * Sample code: Get cloud service network interface.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceNetworkInterface(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .getCloudServiceNetworkInterfaceWithResponse("rg1", "cs1", "TestVMRole_IN_0", "nic1", null, Context.NONE);
    }
}
