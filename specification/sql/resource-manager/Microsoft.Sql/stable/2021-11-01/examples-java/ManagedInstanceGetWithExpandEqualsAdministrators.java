
import com.azure.core.util.Context;

/** Samples for ManagedInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceGetWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: Get managed instance with $expand=administrators/activedirectory.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedInstanceWithExpandAdministratorsActivedirectory(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().getByResourceGroupWithResponse("testrg",
            "testinstance", null, Context.NONE);
    }
}
