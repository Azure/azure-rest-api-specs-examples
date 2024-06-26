
/**
 * Samples for VmmServers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmmServers_Delete_MinimumSet_Gen.
     * json
     */
    /**
     * Sample code: VmmServers_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmmServersDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().delete("rgscvmm", "8", null, com.azure.core.util.Context.NONE);
    }
}
