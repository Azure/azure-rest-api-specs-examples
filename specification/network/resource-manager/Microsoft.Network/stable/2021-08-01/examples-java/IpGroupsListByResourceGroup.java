import com.azure.core.util.Context;

/** Samples for IpGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/IpGroupsListByResourceGroup.json
     */
    /**
     * Sample code: ListByResourceGroup_IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listByResourceGroupIpGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpGroups().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
