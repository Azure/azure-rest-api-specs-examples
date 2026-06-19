
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for NetworkInterfaces UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkInterfaces_UpdateAdministrativeState.json
     */
    /**
     * Sample code: NetworkInterfaces_UpdateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().updateAdministrativeState("example-rg", "example-device", "example-interface",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
