/** Samples for ManagementGroupSubscriptions GetSubscriptionsUnderManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetAllSubscriptionsFromManagementGroup.json
     */
    /**
     * Sample code: GetAllSubscriptionsFromManagementGroup.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void getAllSubscriptionsFromManagementGroup(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager
            .managementGroupSubscriptions()
            .getSubscriptionsUnderManagementGroup("Group", null, com.azure.core.util.Context.NONE);
    }
}
