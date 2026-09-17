
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCrossConnectionInner;
import com.azure.resourcemanager.network.models.ServiceProviderProvisioningState;

/**
 * Samples for ExpressRouteCrossConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteCrossConnectionUpdate.json
     */
    /**
     * Sample code: UpdateExpressRouteCrossConnection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void updateExpressRouteCrossConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().createOrUpdate(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", new ExpressRouteCrossConnectionInner()
                .withServiceProviderProvisioningState(ServiceProviderProvisioningState.NOT_PROVISIONED),
            com.azure.core.util.Context.NONE);
    }
}
