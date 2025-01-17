
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for L2IsolationDomains UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * L2IsolationDomains_updateAdministrativeState_MaximumSet_Gen.json
     */
    /**
     * Sample code: L2IsolationDomains_updateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().updateAdministrativeState("example-rg", "example-l2Domain",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
