
import com.azure.resourcemanager.edgeactions.models.EdgeAction;
import com.azure.resourcemanager.edgeactions.models.SkuTypeUpdate;

/**
 * Samples for EdgeActions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActions_Update.json
     */
    /**
     * Sample code: UpdateEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void updateEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        EdgeAction resource = manager.edgeActions()
            .getByResourceGroupWithResponse("testrg", "edgeAction1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSku(new SkuTypeUpdate().withName("Standard").withTier("Standard")).apply();
    }
}
