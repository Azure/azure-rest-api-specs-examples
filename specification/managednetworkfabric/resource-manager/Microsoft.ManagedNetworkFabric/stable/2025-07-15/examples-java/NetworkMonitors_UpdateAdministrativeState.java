
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for NetworkMonitors UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkMonitors_UpdateAdministrativeState.json
     */
    /**
     * Sample code: NetworkMonitors_UpdateAdministrativeState.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkMonitorsUpdateAdministrativeState(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkMonitors().updateAdministrativeState("example-rg", "example-monitor",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
