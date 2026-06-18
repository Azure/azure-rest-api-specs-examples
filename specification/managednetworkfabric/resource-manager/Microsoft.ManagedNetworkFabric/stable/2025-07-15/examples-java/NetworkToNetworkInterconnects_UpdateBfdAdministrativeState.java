
import com.azure.resourcemanager.managednetworkfabric.models.BfdAdministrativeState;
import com.azure.resourcemanager.managednetworkfabric.models.NniUpdateBfdAdministrativeStateRequest;
import com.azure.resourcemanager.managednetworkfabric.models.RouteType;

/**
 * Samples for NetworkToNetworkInterconnects UpdateBfdAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkToNetworkInterconnects_UpdateBfdAdministrativeState.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_UpdateBfdAdministrativeState.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsUpdateBfdAdministrativeState(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkToNetworkInterconnects().updateBfdAdministrativeState(
            "example-rg", "example-nf", "example-nni", new NniUpdateBfdAdministrativeStateRequest()
                .withRouteType(RouteType.STATIC).withAdministrativeState(BfdAdministrativeState.fromString("Enable")),
            com.azure.core.util.Context.NONE);
    }
}
