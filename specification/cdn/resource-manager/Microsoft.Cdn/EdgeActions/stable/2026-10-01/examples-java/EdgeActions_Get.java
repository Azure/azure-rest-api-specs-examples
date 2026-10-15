
/**
 * Samples for EdgeActions GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActions_Get.json
     */
    /**
     * Sample code: GetEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void getEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActions().getByResourceGroupWithResponse("testrg", "edgeAction1", com.azure.core.util.Context.NONE);
    }
}
