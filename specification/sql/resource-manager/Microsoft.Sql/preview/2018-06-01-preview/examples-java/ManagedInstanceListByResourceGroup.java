import com.azure.core.util.Context;

/** Samples for ManagedInstances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ManagedInstanceListByResourceGroup.json
     */
    /**
     * Sample code: List managed instances by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedInstancesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().listByResourceGroup("Test1", Context.NONE);
    }
}
