
import com.azure.resourcemanager.managednetworkfabric.models.BfdAdministrativeState;
import com.azure.resourcemanager.managednetworkfabric.models.ExternalNetworkRouteType;
import com.azure.resourcemanager.managednetworkfabric.models.ExternalNetworkUpdateBfdAdministrativeStateRequest;

/**
 * Samples for ExternalNetworks UpdateBfdAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/ExternalNetworks_UpdateBfdAdministrativeState.json
     */
    /**
     * Sample code: ExternalNetworks_UpdateBfdAdministrativeState.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksUpdateBfdAdministrativeState(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().updateBfdAdministrativeState("example-rg", "example-externalnetwork", "example-ext",
            new ExternalNetworkUpdateBfdAdministrativeStateRequest().withRouteType(ExternalNetworkRouteType.STATIC)
                .withAdministrativeState(BfdAdministrativeState.fromString("Enable")),
            com.azure.core.util.Context.NONE);
    }
}
