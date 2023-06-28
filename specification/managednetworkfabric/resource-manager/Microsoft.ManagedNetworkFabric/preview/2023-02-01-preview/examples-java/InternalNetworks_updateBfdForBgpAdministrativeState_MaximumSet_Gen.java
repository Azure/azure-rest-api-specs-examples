import com.azure.resourcemanager.managednetworkfabric.models.AdministrativeState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/** Samples for InternalNetworks UpdateBfdForBgpAdministrativeState. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_updateBfdForBgpAdministrativeState_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternalNetworks_updateBfdForBgpAdministrativeState_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksUpdateBfdForBgpAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internalNetworks()
            .updateBfdForBgpAdministrativeState(
                "resourceGroupName",
                "example-l3domain",
                "example-internalnetwork",
                new UpdateAdministrativeState()
                    .withResourceIds(
                        Arrays
                            .asList(
                                "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/example-l3domain/internalNetworks/example-internalnetwork"))
                    .withState(AdministrativeState.ENABLE),
                com.azure.core.util.Context.NONE);
    }
}
