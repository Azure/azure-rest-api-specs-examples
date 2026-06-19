
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for InternalNetworks UpdateStaticRouteBfdAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternalNetworks_UpdateStaticRouteBfdAdministrativeState.json
     */
    /**
     * Sample code: InternalNetworks_UpdateStaticRouteBfdAdministrativeState.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksUpdateStaticRouteBfdAdministrativeState(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internalNetworks().updateStaticRouteBfdAdministrativeState("example-rg", "example-l3domain",
            "example-internalNetwork",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
