
import com.azure.resourcemanager.appcontainers.models.WorkflowArtifacts;
import java.util.Arrays;

/**
 * Samples for LogicApps DeployWorkflowArtifacts.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * LogicApps_DeleteDeployWorkflowArtifacts.json
     */
    /**
     * Sample code: Delete workflow artifacts.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        deleteWorkflowArtifacts(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().deployWorkflowArtifactsWithResponse("testrg123", "testapp2", "testapp2",
            new WorkflowArtifacts().withFilesToDelete(Arrays.asList("test/workflow.json", "test/")),
            com.azure.core.util.Context.NONE);
    }
}
