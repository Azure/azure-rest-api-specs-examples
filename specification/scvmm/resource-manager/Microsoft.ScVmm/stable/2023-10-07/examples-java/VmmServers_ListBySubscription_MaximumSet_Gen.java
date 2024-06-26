
/**
 * Samples for VmmServers List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VmmServers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: VmmServers_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersListBySubscriptionMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().list(com.azure.core.util.Context.NONE);
    }
}
