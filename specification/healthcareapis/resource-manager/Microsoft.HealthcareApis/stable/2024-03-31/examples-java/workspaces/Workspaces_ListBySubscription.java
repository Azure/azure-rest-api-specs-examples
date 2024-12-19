
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/workspaces/
     * Workspaces_ListBySubscription.json
     */
    /**
     * Sample code: Get workspaces by subscription.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        getWorkspacesBySubscription(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}
