import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/** Samples for NetworkFabrics UpdateInfraManagementBfdConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_UpdateInfraManagementBfdConfiguration_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_UpdateInfraManagementBfdConfiguration_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsUpdateInfraManagementBfdConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabrics()
            .updateInfraManagementBfdConfiguration(
                "example-rg",
                "example-fabric",
                new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
                com.azure.core.util.Context.NONE);
    }
}
