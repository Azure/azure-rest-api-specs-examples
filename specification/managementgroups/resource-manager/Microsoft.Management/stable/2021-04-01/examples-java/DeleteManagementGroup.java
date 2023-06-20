/** Samples for ManagementGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/DeleteManagementGroup.json
     */
    /**
     * Sample code: DeleteManagementGroup.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void deleteManagementGroup(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.managementGroups().delete("GroupToDelete", "no-cache", com.azure.core.util.Context.NONE);
    }
}
