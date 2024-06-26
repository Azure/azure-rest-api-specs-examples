
/**
 * Samples for VmmServers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_Get_MaximumSet_Gen.
     * json
     */
    /**
     * Sample code: VmmServers_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().getByResourceGroupWithResponse("rgscvmm", ".", com.azure.core.util.Context.NONE);
    }
}
