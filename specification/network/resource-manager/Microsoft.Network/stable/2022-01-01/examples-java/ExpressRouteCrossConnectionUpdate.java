import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCrossConnectionInner;
import com.azure.resourcemanager.network.models.ServiceProviderProvisioningState;

/** Samples for ExpressRouteCrossConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteCrossConnectionUpdate.json
     */
    /**
     * Sample code: UpdateExpressRouteCrossConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateExpressRouteCrossConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCrossConnections()
            .createOrUpdate(
                "CrossConnection-SiliconValley",
                "<circuitServiceKey>",
                new ExpressRouteCrossConnectionInner()
                    .withServiceProviderProvisioningState(ServiceProviderProvisioningState.NOT_PROVISIONED),
                Context.NONE);
    }
}
