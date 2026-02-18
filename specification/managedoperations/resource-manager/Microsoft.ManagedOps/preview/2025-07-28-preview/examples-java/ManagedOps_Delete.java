
/**
 * Samples for ManagedOps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-28-preview/ManagedOps_Delete.json
     */
    /**
     * Sample code: ManagedOps_Delete.
     * 
     * @param manager Entry point to ManagedOpsManager.
     */
    public static void managedOpsDelete(com.azure.resourcemanager.managedops.ManagedOpsManager manager) {
        manager.managedOps().delete("default", com.azure.core.util.Context.NONE);
    }
}
