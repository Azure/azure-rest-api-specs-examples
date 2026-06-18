
import com.azure.resourcemanager.managednetworkfabric.models.NetworkFabricLockAction;
import com.azure.resourcemanager.managednetworkfabric.models.NetworkFabricLockRequest;
import com.azure.resourcemanager.managednetworkfabric.models.NetworkFabricLockType;

/**
 * Samples for NetworkFabrics LockFabric.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_LockFabric.json
     */
    /**
     * Sample code: NetworkFabrics_LockFabric_MaximumSet.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsLockFabricMaximumSet(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().lockFabric(
            "example-rg", "example-networkFabric", new NetworkFabricLockRequest()
                .withLockType(NetworkFabricLockType.ADMINISTRATIVE).withAction(NetworkFabricLockAction.LOCK),
            com.azure.core.util.Context.NONE);
    }
}
