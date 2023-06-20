import com.azure.resourcemanager.managementgroups.models.CreateOrUpdateSettingsRequest;

/** Samples for HierarchySettingsOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutHierarchySettings.json
     */
    /**
     * Sample code: GetGroupSettings.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getGroupSettings(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager
            .hierarchySettingsOperations()
            .createOrUpdateWithResponse(
                "root",
                new CreateOrUpdateSettingsRequest()
                    .withRequireAuthorizationForGroupCreation(true)
                    .withDefaultManagementGroup("/providers/Microsoft.Management/managementGroups/DefaultGroup"),
                com.azure.core.util.Context.NONE);
    }
}
