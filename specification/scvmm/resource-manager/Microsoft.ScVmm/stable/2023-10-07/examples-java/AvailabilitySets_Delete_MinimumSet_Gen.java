
/**
 * Samples for AvailabilitySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * AvailabilitySets_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void availabilitySetsDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.availabilitySets().delete("rgscvmm", "6", null, com.azure.core.util.Context.NONE);
    }
}
