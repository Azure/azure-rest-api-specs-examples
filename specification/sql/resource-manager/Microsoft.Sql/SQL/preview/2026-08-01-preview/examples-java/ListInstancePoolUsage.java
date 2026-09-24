
/**
 * Samples for Usages ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListInstancePoolUsage.json
     */
    /**
     * Sample code: List instance pool usages.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listInstancePoolUsages(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getUsages().listByInstancePool("group1", "testIP", null,
            com.azure.core.util.Context.NONE);
    }
}
