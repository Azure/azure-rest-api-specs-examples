
/**
 * Samples for DedicatedHostGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * dedicatedHostExamples/DedicatedHostGroup_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_Delete_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostGroupDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHostGroups().deleteWithResponse("rgcompute", "a",
            com.azure.core.util.Context.NONE);
    }
}
