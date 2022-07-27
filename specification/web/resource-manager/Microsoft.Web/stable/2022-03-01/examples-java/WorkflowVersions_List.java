import com.azure.core.util.Context;

/** Samples for WorkflowVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowVersions_List.json
     */
    /**
     * Sample code: List a workflows versions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAWorkflowsVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWorkflowVersions()
            .list("test-resource-group", "test-name", "test-workflow", null, Context.NONE);
    }
}
