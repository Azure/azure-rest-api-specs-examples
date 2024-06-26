
/**
 * Samples for GuestAgents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GuestAgents_Get_MaximumSet_Gen.
     * json
     */
    /**
     * Sample code: GuestAgents_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void guestAgentsGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.guestAgents().getWithResponse("gtgclehcbsyave", com.azure.core.util.Context.NONE);
    }
}
