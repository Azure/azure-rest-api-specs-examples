import com.azure.core.util.Context;

/** Samples for DedicatedHostGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/dedicatedHostExamples/DedicatedHostGroups_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroups_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostGroupsDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHostGroups()
            .deleteWithResponse("rgcompute", "aaaa", Context.NONE);
    }
}
