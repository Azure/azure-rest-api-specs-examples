import com.azure.core.util.Context;

/** Samples for VirtualWans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualWANList.json
     */
    /**
     * Sample code: VirtualWANList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualWANList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualWans().list(Context.NONE);
    }
}
