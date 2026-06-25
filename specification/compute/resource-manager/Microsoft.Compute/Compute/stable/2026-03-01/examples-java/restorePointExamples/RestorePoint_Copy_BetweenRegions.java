
import com.azure.resourcemanager.compute.fluent.models.RestorePointInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;

/**
 * Samples for RestorePoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePoint_Copy_BetweenRegions.json
     */
    /**
     * Sample code: Copy a restore point to a different region.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void copyARestorePointToADifferentRegion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePoints().create("myResourceGroup", "rpcName", "rpName",
            new RestorePointInner().withSourceRestorePoint(new ApiEntityReference().withId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName")),
            com.azure.core.util.Context.NONE);
    }
}
