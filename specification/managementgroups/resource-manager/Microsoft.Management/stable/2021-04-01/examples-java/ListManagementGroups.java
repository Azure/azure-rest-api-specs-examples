/** Samples for ManagementGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/ListManagementGroups.json
     */
    /**
     * Sample code: ListManagementGroups.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void listManagementGroups(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.managementGroups().list("no-cache", null, com.azure.core.util.Context.NONE);
    }
}
