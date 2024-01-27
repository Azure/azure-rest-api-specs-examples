
/** Samples for ClassicAdministrators List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/
     * GetClassicAdministrators.json
     */
    /**
     * Sample code: List classic administrators.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listClassicAdministrators(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getClassicAdministrators()
            .list(com.azure.core.util.Context.NONE);
    }
}
