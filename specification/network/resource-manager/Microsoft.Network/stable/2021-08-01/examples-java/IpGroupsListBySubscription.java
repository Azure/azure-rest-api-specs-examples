import com.azure.core.util.Context;

/** Samples for IpGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/IpGroupsListBySubscription.json
     */
    /**
     * Sample code: List_IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listIpGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpGroups().list(Context.NONE);
    }
}
