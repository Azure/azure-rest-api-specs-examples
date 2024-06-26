
/**
 * Samples for VmmServers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_Get_MinimumSet_Gen.
     * json
     */
    /**
     * Sample code: VmmServers_Get_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersGetMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().getByResourceGroupWithResponse("rgscvmm", "D", com.azure.core.util.Context.NONE);
    }
}
