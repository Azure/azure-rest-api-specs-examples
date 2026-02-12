
/**
 * Samples for DisconnectedOperations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DisconnectedOperations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DisconnectedOperations_Get.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void disconnectedOperationsGet(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.disconnectedOperations().getByResourceGroupWithResponse("rgdisconnectedoperations", "demo-resource",
            com.azure.core.util.Context.NONE);
    }
}
