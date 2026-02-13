
/**
 * Samples for DisconnectedOperations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DisconnectedOperations_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DisconnectedOperations_ListByResourceGroup.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void disconnectedOperationsListByResourceGroup(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.disconnectedOperations().listByResourceGroup("rgdisconnectedoperations",
            com.azure.core.util.Context.NONE);
    }
}
