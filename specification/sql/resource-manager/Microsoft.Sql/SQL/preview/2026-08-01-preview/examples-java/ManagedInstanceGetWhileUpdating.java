
/**
 * Samples for ManagedInstances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceGetWhileUpdating.json
     */
    /**
     * Sample code: Get managed instance while resource is updating.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getManagedInstanceWhileResourceIsUpdating(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().getByResourceGroupWithResponse("testrg", "testinstance", null,
            com.azure.core.util.Context.NONE);
    }
}
