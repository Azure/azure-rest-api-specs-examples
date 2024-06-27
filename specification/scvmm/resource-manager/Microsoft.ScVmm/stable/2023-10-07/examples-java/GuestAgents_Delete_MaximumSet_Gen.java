
/**
 * Samples for GuestAgents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GuestAgents_Delete_MaximumSet_Gen
     * .json
     */
    /**
     * Sample code: GuestAgents_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void guestAgentsDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.guestAgents().deleteWithResponse("gtgclehcbsyave", com.azure.core.util.Context.NONE);
    }
}
