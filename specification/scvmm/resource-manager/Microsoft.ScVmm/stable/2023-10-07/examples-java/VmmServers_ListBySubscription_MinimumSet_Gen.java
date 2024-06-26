
/**
 * Samples for VmmServers List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VmmServers_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: VmmServers_ListBySubscription_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersListBySubscriptionMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().list(com.azure.core.util.Context.NONE);
    }
}
