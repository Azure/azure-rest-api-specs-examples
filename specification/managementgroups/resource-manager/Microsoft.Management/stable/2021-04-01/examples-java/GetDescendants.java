
/**
 * Samples for ManagementGroups GetDescendants.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetDescendants.
     * json
     */
    /**
     * Sample code: GetDescendants.
     * 
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getDescendants(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.managementGroups().getDescendants("20000000-0000-0000-0000-000000000000", null, null,
            com.azure.core.util.Context.NONE);
    }
}
