
/**
 * Samples for ManagedInstanceAzureADOnlyAuthentications ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceAzureADOnlyAuthListByInstance.json
     */
    /**
     * Sample code: Gets a list of Azure Active Directory only authentication object.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfAzureActiveDirectoryOnlyAuthenticationObject(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAzureADOnlyAuthentications()
            .listByInstance("Default-SQL-SouthEastAsia", "managedInstance", com.azure.core.util.Context.NONE);
    }
}
