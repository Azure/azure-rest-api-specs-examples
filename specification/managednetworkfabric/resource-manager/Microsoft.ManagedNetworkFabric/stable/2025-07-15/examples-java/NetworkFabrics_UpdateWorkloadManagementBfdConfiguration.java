
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for NetworkFabrics UpdateWorkloadManagementBfdConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_UpdateWorkloadManagementBfdConfiguration.json
     */
    /**
     * Sample code: NetworkFabrics_UpdateWorkloadManagementBfdConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsUpdateWorkloadManagementBfdConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().updateWorkloadManagementBfdConfiguration("example-rg", "example-fabric",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
