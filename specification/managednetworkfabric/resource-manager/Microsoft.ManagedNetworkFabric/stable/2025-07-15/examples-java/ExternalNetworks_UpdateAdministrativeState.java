
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for ExternalNetworks UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/ExternalNetworks_UpdateAdministrativeState.json
     */
    /**
     * Sample code: ExternalNetworks_UpdateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().updateAdministrativeState("example-rg", "example-externalnetwork", "example-ext",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
