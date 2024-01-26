
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.WorkloadClassifierInner;

/** Samples for WorkloadClassifiers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateWorkloadClassifierMax.
     * json
     */
    /**
     * Sample code: Create a workload group with all properties specified.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAWorkloadGroupWithAllPropertiesSpecified(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadClassifiers()
            .createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup",
                "wlm_workloadclassifier", new WorkloadClassifierInner().withMemberName("dbo").withLabel("test_label")
                    .withContext("test_context").withStartTime("12:00").withEndTime("14:00").withImportance("high"),
                Context.NONE);
    }
}
