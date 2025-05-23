
/**
 * Samples for ManagementGroupSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/
     * RemoveManagementGroupSubscription.json
     */
    /**
     * Sample code: DeleteSubscriptionFromManagementGroup.
     * 
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void deleteSubscriptionFromManagementGroup(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.managementGroupSubscriptions().deleteWithResponse("Group", "728bcbe4-8d56-4510-86c2-4921b8beefbc",
            "no-cache", com.azure.core.util.Context.NONE);
    }
}
