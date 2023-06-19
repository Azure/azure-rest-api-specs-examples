/** Samples for HierarchySettingsOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/DeleteHierarchySettings.json
     */
    /**
     * Sample code: GetGroupSettings.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getGroupSettings(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.hierarchySettingsOperations().deleteWithResponse("root", com.azure.core.util.Context.NONE);
    }
}
