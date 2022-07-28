import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.WorkloadGroupInner;

/** Samples for WorkloadGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/CreateOrUpdateWorkloadGroupMin.json
     */
    /**
     * Sample code: Create a workload group with the required properties specified.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAWorkloadGroupWithTheRequiredPropertiesSpecified(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getWorkloadGroups()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "testdb",
                "smallrc",
                new WorkloadGroupInner()
                    .withMinResourcePercent(0)
                    .withMaxResourcePercent(100)
                    .withMinResourcePercentPerRequest(3.0),
                Context.NONE);
    }
}
