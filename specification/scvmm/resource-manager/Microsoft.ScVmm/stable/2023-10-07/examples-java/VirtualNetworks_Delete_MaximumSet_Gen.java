
import com.azure.resourcemanager.scvmm.models.ForceDelete;

/**
 * Samples for VirtualNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualNetworks_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualNetworks_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualNetworksDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().delete("rgscvmm", ".", ForceDelete.TRUE, com.azure.core.util.Context.NONE);
    }
}
