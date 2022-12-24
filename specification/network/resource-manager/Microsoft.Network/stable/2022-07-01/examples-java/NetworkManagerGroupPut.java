import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.NetworkGroupInner;

/** Samples for NetworkGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkManagerGroupPut.json
     */
    /**
     * Sample code: NetworkGroupsPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkGroupsPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkGroups()
            .createOrUpdateWithResponse(
                "rg1",
                "testNetworkManager",
                "testNetworkGroup",
                new NetworkGroupInner().withDescription("A sample group"),
                null,
                Context.NONE);
    }
}
