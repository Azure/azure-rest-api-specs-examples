
/**
 * Samples for DisconnectedOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DisconnectedOperations_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: DisconnectedOperations_ListBySubscription.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void disconnectedOperationsListBySubscription(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.disconnectedOperations().list(com.azure.core.util.Context.NONE);
    }
}
