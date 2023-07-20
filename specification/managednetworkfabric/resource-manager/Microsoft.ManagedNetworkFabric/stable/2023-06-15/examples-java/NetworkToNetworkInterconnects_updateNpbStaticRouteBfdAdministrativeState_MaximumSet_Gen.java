import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/** Samples for NetworkToNetworkInterconnects UpdateNpbStaticRouteBfdAdministrativeState. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_updateNpbStaticRouteBfdAdministrativeState_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkToNetworkInterconnects_updateNpbStaticRouteBfdAdministrativeState_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkToNetworkInterconnectsUpdateNpbStaticRouteBfdAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkToNetworkInterconnects()
            .updateNpbStaticRouteBfdAdministrativeState(
                "example-rg",
                "example-fabric",
                "example-nni",
                new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
                com.azure.core.util.Context.NONE);
    }
}
