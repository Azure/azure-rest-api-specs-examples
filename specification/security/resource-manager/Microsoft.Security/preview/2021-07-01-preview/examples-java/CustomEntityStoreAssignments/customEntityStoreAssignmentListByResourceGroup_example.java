
/**
 * Samples for CustomEntityStoreAssignments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/
     * CustomEntityStoreAssignments/customEntityStoreAssignmentListByResourceGroup_example.json
     */
    /**
     * Sample code: List custom entity store assignments in a subscription and a resource group.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listCustomEntityStoreAssignmentsInASubscriptionAndAResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customEntityStoreAssignments().listByResourceGroup("TestResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
