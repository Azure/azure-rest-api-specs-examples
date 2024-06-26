
/**
 * Samples for AvailabilitySets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * AvailabilitySets_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void availabilitySetsGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.availabilitySets().getByResourceGroupWithResponse("rgscvmm", "-", com.azure.core.util.Context.NONE);
    }
}
