
import com.azure.resourcemanager.managedops.models.ManagedOp;

/**
 * Samples for ManagedOps Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-28-preview/ManagedOps_Update.json
     */
    /**
     * Sample code: ManagedOps_Update.
     * 
     * @param manager Entry point to ManagedOpsManager.
     */
    public static void managedOpsUpdate(com.azure.resourcemanager.managedops.ManagedOpsManager manager) {
        ManagedOp resource
            = manager.managedOps().getWithResponse("default", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
