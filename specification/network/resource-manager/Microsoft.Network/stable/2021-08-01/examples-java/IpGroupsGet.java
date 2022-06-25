import com.azure.core.util.Context;

/** Samples for IpGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/IpGroupsGet.json
     */
    /**
     * Sample code: Get_IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getIpGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getIpGroups()
            .getByResourceGroupWithResponse("myResourceGroup", "ipGroups1", null, Context.NONE);
    }
}
