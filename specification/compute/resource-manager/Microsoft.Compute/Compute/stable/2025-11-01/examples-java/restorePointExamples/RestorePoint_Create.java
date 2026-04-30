
import com.azure.resourcemanager.compute.fluent.models.RestorePointInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import java.util.Arrays;

/**
 * Samples for RestorePoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePoint_Create.json
     */
    /**
     * Sample code: Create a restore point.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createARestorePoint(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePoints().create("myResourceGroup", "rpcName", "rpName",
            new RestorePointInner().withExcludeDisks(Arrays.asList(new ApiEntityReference().withId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123")))
                .withInstantAccessDurationMinutes(120),
            com.azure.core.util.Context.NONE);
    }
}
