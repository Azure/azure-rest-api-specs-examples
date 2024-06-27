
import com.azure.resourcemanager.scvmm.models.ForceDelete;

/**
 * Samples for Clouds Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Clouds_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clouds_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void cloudsDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().delete("rgscvmm", "-", ForceDelete.TRUE, com.azure.core.util.Context.NONE);
    }
}
