
import com.azure.core.util.Context;

/** Samples for DistributedAvailabilityGroups ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * DistributedAvailabilityGroupsListByInstance.json
     */
    /**
     * Sample code: Lists all distributed availability groups in instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listsAllDistributedAvailabilityGroupsInInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDistributedAvailabilityGroups().listByInstance("testrg",
            "testcl", Context.NONE);
    }
}
