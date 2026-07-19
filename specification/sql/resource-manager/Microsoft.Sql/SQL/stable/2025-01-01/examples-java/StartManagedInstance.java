
/**
 * Samples for ManagedInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/StartManagedInstance.json
     */
    /**
     * Sample code: Starts the managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void startsTheManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().start("startrg", "mitostart", com.azure.core.util.Context.NONE);
    }
}
