
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for ExternalNetworks UpdateStaticRouteBfdAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/ExternalNetworks_UpdateStaticRouteBfdAdministrativeState.json
     */
    /**
     * Sample code: ExternalNetworks_UpdateStaticRouteBfdAdministrativeState.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksUpdateStaticRouteBfdAdministrativeState(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().updateStaticRouteBfdAdministrativeState("example-rg", "example-l3domain",
            "example-externalnetwork",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
