
/**
 * Samples for Usages ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListInstancePoolUsageExpanded.json
     */
    /**
     * Sample code: List instance pool usages expanded with children.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listInstancePoolUsagesExpandedWithChildren(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getUsages().listByInstancePool("group1", "testIP", true,
            com.azure.core.util.Context.NONE);
    }
}
