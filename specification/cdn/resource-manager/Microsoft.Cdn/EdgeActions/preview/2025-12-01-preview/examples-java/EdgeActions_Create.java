
import com.azure.resourcemanager.edgeactions.models.SkuType;

/**
 * Samples for EdgeActions Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActions_Create.json
     */
    /**
     * Sample code: CreateEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void createEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActions().define("edgeAction1").withRegion("global").withExistingResourceGroup("testrg")
            .withSku(new SkuType().withName("Standard").withTier("Standard")).create();
    }
}
