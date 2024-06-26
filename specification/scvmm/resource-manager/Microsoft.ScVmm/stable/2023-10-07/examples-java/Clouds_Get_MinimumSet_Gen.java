
/**
 * Samples for Clouds GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Clouds_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: Clouds_Get_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void cloudsGetMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().getByResourceGroupWithResponse("rgscvmm", "i", com.azure.core.util.Context.NONE);
    }
}
