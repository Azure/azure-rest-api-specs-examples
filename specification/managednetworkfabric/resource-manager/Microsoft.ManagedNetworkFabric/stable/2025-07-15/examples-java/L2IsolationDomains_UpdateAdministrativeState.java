
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for L2IsolationDomains UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/L2IsolationDomains_UpdateAdministrativeState.json
     */
    /**
     * Sample code: L2IsolationDomains_UpdateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().updateAdministrativeState("example-rg", "example-l2domain",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
