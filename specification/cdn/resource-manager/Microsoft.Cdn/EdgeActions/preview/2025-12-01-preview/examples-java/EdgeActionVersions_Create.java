
import com.azure.resourcemanager.edgeactions.fluent.models.EdgeActionVersionPropertiesInner;
import com.azure.resourcemanager.edgeactions.models.EdgeActionIsDefaultVersion;
import com.azure.resourcemanager.edgeactions.models.EdgeActionVersionDeploymentType;

/**
 * Samples for EdgeActionVersions Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionVersions_Create.json
     */
    /**
     * Sample code: CreateEdgeActionVersion.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void createEdgeActionVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().define("version2").withRegion("global")
            .withExistingEdgeAction("testrg", "edgeAction1")
            .withProperties(
                new EdgeActionVersionPropertiesInner().withDeploymentType(EdgeActionVersionDeploymentType.ZIP)
                    .withIsDefaultVersion(EdgeActionIsDefaultVersion.TRUE))
            .create();
    }
}
