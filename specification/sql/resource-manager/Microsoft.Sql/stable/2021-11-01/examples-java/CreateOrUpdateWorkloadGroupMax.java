
import com.azure.resourcemanager.sql.fluent.models.WorkloadGroupInner;

/**
 * Samples for WorkloadGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateWorkloadGroupMax.json
     */
    /**
     * Sample code: Create a workload group with all properties specified.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAWorkloadGroupWithAllPropertiesSpecified(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadGroups().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", "smallrc",
            new WorkloadGroupInner().withMinResourcePercent(0).withMaxResourcePercent(100)
                .withMinResourcePercentPerRequest(3.0).withMaxResourcePercentPerRequest(3.0D).withImportance("normal")
                .withQueryExecutionTimeout(0),
            com.azure.core.util.Context.NONE);
    }
}
