
/**
 * Samples for VirtualNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualNetworks_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualNetworks_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualNetworksDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().delete("rgscvmm", "1", null, com.azure.core.util.Context.NONE);
    }
}
