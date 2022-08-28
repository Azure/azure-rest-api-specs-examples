import com.azure.core.util.Context;

/** Samples for DscpConfiguration List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/DscpConfigurationListAll.json
     */
    /**
     * Sample code: List all network interfaces.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDscpConfigurations().list(Context.NONE);
    }
}
