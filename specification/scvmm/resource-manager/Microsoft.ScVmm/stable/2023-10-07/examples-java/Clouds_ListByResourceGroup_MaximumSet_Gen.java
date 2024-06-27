
/**
 * Samples for Clouds ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * Clouds_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clouds_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void cloudsListByResourceGroupMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().listByResourceGroup("rgscvmm", com.azure.core.util.Context.NONE);
    }
}
