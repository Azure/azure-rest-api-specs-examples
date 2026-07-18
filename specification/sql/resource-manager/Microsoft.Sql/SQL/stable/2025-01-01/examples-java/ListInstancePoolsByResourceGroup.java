
/**
 * Samples for InstancePools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListInstancePoolsByResourceGroup.json
     */
    /**
     * Sample code: List instance pools by resource group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listInstancePoolsByResourceGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstancePools().listByResourceGroup("group1", com.azure.core.util.Context.NONE);
    }
}
