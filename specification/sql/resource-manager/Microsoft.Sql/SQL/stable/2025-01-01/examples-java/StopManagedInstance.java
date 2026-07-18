
/**
 * Samples for ManagedInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/StopManagedInstance.json
     */
    /**
     * Sample code: Stops the managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void stopsTheManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().stop("stoprg", "mitostop", com.azure.core.util.Context.NONE);
    }
}
