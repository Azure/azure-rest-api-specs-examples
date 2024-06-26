
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualNetworks_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualNetworks_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualNetworksGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().getByResourceGroupWithResponse("rgscvmm", "2", com.azure.core.util.Context.NONE);
    }
}
