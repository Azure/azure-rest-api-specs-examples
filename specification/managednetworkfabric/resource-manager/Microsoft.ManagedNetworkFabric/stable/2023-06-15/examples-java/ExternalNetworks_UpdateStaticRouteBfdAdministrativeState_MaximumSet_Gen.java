import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/** Samples for ExternalNetworks UpdateStaticRouteBfdAdministrativeState. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_UpdateStaticRouteBfdAdministrativeState_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalNetworks_UpdateStaticRouteBfdAdministrativeState_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksUpdateStaticRouteBfdAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .externalNetworks()
            .updateStaticRouteBfdAdministrativeState(
                "example-rg",
                "example-l3domain",
                "example-externalnetwork",
                new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
                com.azure.core.util.Context.NONE);
    }
}
