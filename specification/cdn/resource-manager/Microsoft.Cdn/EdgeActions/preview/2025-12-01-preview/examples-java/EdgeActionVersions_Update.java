
import com.azure.resourcemanager.edgeactions.models.EdgeActionVersion;
import com.azure.resourcemanager.edgeactions.models.EdgeActionVersionDeploymentType;
import com.azure.resourcemanager.edgeactions.models.EdgeActionVersionUpdateProperties;

/**
 * Samples for EdgeActionVersions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionVersions_Update.json
     */
    /**
     * Sample code: UpdateEdgeActionVersion.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void updateEdgeActionVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        EdgeActionVersion resource = manager.edgeActionVersions()
            .getWithResponse("testrg", "edgeAction1", "version1", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(
                new EdgeActionVersionUpdateProperties().withDeploymentType(EdgeActionVersionDeploymentType.OTHERS))
            .apply();
    }
}
