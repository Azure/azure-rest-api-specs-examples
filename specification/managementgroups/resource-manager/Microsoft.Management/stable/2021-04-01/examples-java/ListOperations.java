/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/ListOperations.json
     */
    /**
     * Sample code: List Operations.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void listOperations(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
