import com.azure.core.util.Context;

/** Samples for NetworkGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerGroupList.json
     */
    /**
     * Sample code: NetworkGroupsList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkGroupsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkGroups()
            .list("rg1", "testNetworkManager", null, null, Context.NONE);
    }
}
