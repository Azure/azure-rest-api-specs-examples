
import com.azure.resourcemanager.edgeactions.fluent.models.VersionCodeInner;

/**
 * Samples for EdgeActionVersions DeployVersionCode.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionVersions_DeployVersionCode.json
     */
    /**
     * Sample code: DeployVersionCode.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void deployVersionCode(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().deployVersionCode("testrg", "edgeAction1", "version2",
            new VersionCodeInner().withContent("UEsDBBQAAAAIAI1NzkQAAAAABQAAAA==").withName("zippedFile"),
            com.azure.core.util.Context.NONE);
    }
}
