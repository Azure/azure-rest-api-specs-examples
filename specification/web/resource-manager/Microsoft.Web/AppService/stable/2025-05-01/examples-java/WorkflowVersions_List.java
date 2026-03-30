
/**
 * Samples for WorkflowVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowVersions_List.json
     */
    /**
     * Sample code: List a workflows versions.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAWorkflowsVersions(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowVersions().list("test-resource-group", "test-name", "test-workflow", null,
            com.azure.core.util.Context.NONE);
    }
}
