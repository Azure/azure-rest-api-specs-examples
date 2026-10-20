
import com.azure.resourcemanager.edgeactions.models.EdgeActionIsDefaultVersion;
import com.azure.resourcemanager.edgeactions.models.EdgeActionVersionDeploymentType;
import com.azure.resourcemanager.edgeactions.models.EdgeActionVersionProperties;

/**
 * Samples for EdgeActionVersions Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionVersions_Create.json
     */
    /**
     * Sample code: CreateEdgeActionVersion.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void createEdgeActionVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().define("version2").withRegion("global")
            .withExistingEdgeAction("testrg", "edgeAction1")
            .withProperties(new EdgeActionVersionProperties().withDeploymentType(EdgeActionVersionDeploymentType.ZIP)
                .withIsDefaultVersion(EdgeActionIsDefaultVersion.TRUE))
            .create();
    }
}
