import com.azure.core.util.Context;

/** Samples for ExpressRouteCrossConnectionPeerings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteCrossConnectionBgpPeeringList.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionBgpPeeringList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteCrossConnectionBgpPeeringList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCrossConnectionPeerings()
            .list("CrossConnection-SiliconValley", "<circuitServiceKey>", Context.NONE);
    }
}
