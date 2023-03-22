/** Samples for CustomEntityStoreAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentListBySubscription_example.json
     */
    /**
     * Sample code: List custom entity store assignments in a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listCustomEntityStoreAssignmentsInASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customEntityStoreAssignments().list(com.azure.core.util.Context.NONE);
    }
}
