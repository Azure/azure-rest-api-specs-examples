import com.azure.resourcemanager.managementgroups.models.ManagementGroupExpandType;

/** Samples for ManagementGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithAncestors.json
     */
    /**
     * Sample code: GetManagementGroupWithAncestors.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getManagementGroupWithAncestors(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager
            .managementGroups()
            .getWithResponse(
                "20000000-0001-0000-0000-00000000000",
                ManagementGroupExpandType.ANCESTORS,
                null,
                null,
                "no-cache",
                com.azure.core.util.Context.NONE);
    }
}
