
/**
 * Samples for AvailabilitySets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * AvailabilitySets_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Get_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void availabilitySetsGetMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.availabilitySets().getByResourceGroupWithResponse("rgscvmm", "V", com.azure.core.util.Context.NONE);
    }
}
