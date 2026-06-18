
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for InternalNetworks UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternalNetworks_UpdateAdministrativeState.json
     */
    /**
     * Sample code: InternalNetworks_UpdateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internalNetworks().updateAdministrativeState("example-rg", "example-l3isd", "example-internalnetwork",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
