
/**
 * Samples for VmmServers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VmmServers_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: VmmServers_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersListByResourceGroupMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().listByResourceGroup("rgscvmm", com.azure.core.util.Context.NONE);
    }
}
