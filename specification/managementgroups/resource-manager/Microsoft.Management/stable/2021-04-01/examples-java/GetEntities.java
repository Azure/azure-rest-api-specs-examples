/** Samples for Entities List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetEntities.json
     */
    /**
     * Sample code: GetEntities.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getEntities(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.entities().list(null, null, null, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
