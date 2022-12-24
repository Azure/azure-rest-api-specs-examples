import com.azure.core.util.Context;

/** Samples for VirtualHubs GetEffectiveVirtualHubRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/EffectiveRoutesListForVirtualHub.json
     */
    /**
     * Sample code: Effective Routes for the Virtual Hub.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void effectiveRoutesForTheVirtualHub(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubs()
            .getEffectiveVirtualHubRoutes("rg1", "virtualHub1", null, Context.NONE);
    }
}
