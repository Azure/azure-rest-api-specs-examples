
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualNetworks_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualNetworks_Get_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualNetworksGetMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().getByResourceGroupWithResponse("rgscvmm", "-", com.azure.core.util.Context.NONE);
    }
}
