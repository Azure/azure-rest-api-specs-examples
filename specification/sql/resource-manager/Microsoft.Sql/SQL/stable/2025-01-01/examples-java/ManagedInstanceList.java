
/**
 * Samples for ManagedInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceList.json
     */
    /**
     * Sample code: List managed instances.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedInstances(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().list(null, com.azure.core.util.Context.NONE);
    }
}
