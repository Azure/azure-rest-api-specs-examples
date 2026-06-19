
import com.azure.resourcemanager.managednetworkfabric.models.BgpAdministrativeState;
import com.azure.resourcemanager.managednetworkfabric.models.InternalNetworkUpdateBgpAdministrativeStateRequest;

/**
 * Samples for InternalNetworks UpdateBgpAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternalNetworks_UpdateBgpAdministrativeState.json
     */
    /**
     * Sample code: InternalNetworks_UpdateBgpAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksUpdateBgpAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internalNetworks().updateBgpAdministrativeState("example-rg", "example-l3isd",
            "example-internalnetwork",
            new InternalNetworkUpdateBgpAdministrativeStateRequest().withNeighborAddress("10.10.10.10")
                .withAdministrativeState(BgpAdministrativeState.fromString("Enable")),
            com.azure.core.util.Context.NONE);
    }
}
