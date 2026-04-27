
/**
 * Samples for ServiceGroups ListAncestors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01-preview/ServiceGroup_ListAncestors.json
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
