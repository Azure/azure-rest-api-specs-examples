
import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateAdministrativeState;
import java.util.Arrays;

/**
 * Samples for AccessControlLists UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/AccessControlLists_UpdateAdministrativeState.json
     */
    /**
     * Sample code: AccessControlLists_UpdateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().updateAdministrativeState("example-rg", "example-acl",
            new UpdateAdministrativeState().withResourceIds(Arrays.asList("")).withState(EnableDisableState.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
