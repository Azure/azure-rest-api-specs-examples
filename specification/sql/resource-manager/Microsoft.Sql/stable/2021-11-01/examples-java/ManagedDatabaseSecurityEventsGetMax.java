
/**
 * Samples for ManagedDatabaseSecurityEvents ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSecurityEventsGetMax.
     * json
     */
    /**
     * Sample code: Get the managed database's security events with maximal parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheManagedDatabaseSSecurityEventsWithMaximalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSecurityEvents().listByDatabase("testrg",
            "testcl", "database1", "ShowServerRecords eq true", 0L, 1L,
            "eyJCbG9iTmFtZURhdGVUaW1lIjoiXC9EYXRlKDE1MTIyODg4MTIwMTArMDIwMClcLyIsIkJsb2JOYW1lUm9sbG92ZXJJbmRleCI6IjAiLCJFbmREYXRlIjoiXC9EYXRlKDE1MTI0NjYyMDA1MjkpXC8iLCJJc1NraXBUb2tlblNldCI6ZmFsc2UsIklzVjJCbG9iVGltZUZvcm1hdCI6dHJ1ZSwiU2hvd1NlcnZlclJlY29yZHMiOmZhbHNlLCJTa2lwVmFsdWUiOjAsIlRha2VWYWx1ZSI6MTB9",
            com.azure.core.util.Context.NONE);
    }
}
