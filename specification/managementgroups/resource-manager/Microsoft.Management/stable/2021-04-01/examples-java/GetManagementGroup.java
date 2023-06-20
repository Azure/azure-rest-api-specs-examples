/** Samples for ManagementGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroup.json
     */
    /**
     * Sample code: GetManagementGroup.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getManagementGroup(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager
            .managementGroups()
            .getWithResponse(
                "20000000-0001-0000-0000-000000000000", null, null, null, "no-cache", com.azure.core.util.Context.NONE);
    }
}
