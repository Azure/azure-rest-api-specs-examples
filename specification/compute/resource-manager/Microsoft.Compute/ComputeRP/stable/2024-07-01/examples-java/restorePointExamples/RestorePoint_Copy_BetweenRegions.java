
import com.azure.resourcemanager.compute.fluent.models.RestorePointInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;

/**
 * Samples for RestorePoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * restorePointExamples/RestorePoint_Copy_BetweenRegions.json
     */
    /**
     * Sample code: Copy a restore point to a different region.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void copyARestorePointToADifferentRegion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getRestorePoints().create("myResourceGroup", "rpcName",
            "rpName",
            new RestorePointInner().withSourceRestorePoint(new ApiEntityReference().withId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName")),
            com.azure.core.util.Context.NONE);
    }
}
