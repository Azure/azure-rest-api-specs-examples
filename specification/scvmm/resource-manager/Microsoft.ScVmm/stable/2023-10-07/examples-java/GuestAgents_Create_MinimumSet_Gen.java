
import com.azure.resourcemanager.scvmm.fluent.models.GuestAgentInner;

/**
 * Samples for GuestAgents Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GuestAgents_Create_MinimumSet_Gen
     * .json
     */
    /**
     * Sample code: GuestAgents_Create_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void guestAgentsCreateMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.guestAgents().create("gtgclehcbsyave", new GuestAgentInner(), com.azure.core.util.Context.NONE);
    }
}
