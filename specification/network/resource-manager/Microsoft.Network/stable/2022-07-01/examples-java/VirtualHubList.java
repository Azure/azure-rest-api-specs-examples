import com.azure.core.util.Context;

/** Samples for VirtualHubs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualHubList.json
     */
    /**
     * Sample code: VirtualHubList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().list(Context.NONE);
    }
}
