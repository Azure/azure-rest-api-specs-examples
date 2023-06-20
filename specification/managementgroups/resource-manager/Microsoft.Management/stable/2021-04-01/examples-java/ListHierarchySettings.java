/** Samples for HierarchySettingsOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/ListHierarchySettings.json
     */
    /**
     * Sample code: ListGroupSettings.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void listGroupSettings(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.hierarchySettingsOperations().listWithResponse("root", com.azure.core.util.Context.NONE);
    }
}
