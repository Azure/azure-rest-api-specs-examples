import com.azure.core.util.Context;

/** Samples for NetworkInterfaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkInterfaceDelete.json
     */
    /**
     * Sample code: Delete network interface.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkInterface(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces().delete("rg1", "test-nic", Context.NONE);
    }
}
