
/**
 * Samples for ServiceGroups ListAncestors.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/management/resource-manager/Microsoft.Management/ServiceGroups/preview/2024-02-01-preview/examples/
     * ServiceGroup_ListAncestors.json
     */
    /**
     * Sample code: ListServiceGroupAncestors.
     * 
     * @param manager Entry point to ServiceGroupsManager.
     */
    public static void listServiceGroupAncestors(com.azure.resourcemanager.servicegroups.ServiceGroupsManager manager) {
        manager.serviceGroups().listAncestorsWithResponse("20000000-0001-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
