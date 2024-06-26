
/**
 * Samples for Clouds Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Clouds_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Clouds_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void cloudsDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().delete("rgscvmm", "1", null, com.azure.core.util.Context.NONE);
    }
}
