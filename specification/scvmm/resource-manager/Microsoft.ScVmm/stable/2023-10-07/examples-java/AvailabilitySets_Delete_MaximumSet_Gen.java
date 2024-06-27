
import com.azure.resourcemanager.scvmm.models.ForceDelete;

/**
 * Samples for AvailabilitySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * AvailabilitySets_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void availabilitySetsDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.availabilitySets().delete("rgscvmm", "_", ForceDelete.TRUE, com.azure.core.util.Context.NONE);
    }
}
