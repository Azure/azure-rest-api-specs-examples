import com.azure.core.util.Context;

/** Samples for Workspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_ListBySubscription.json
     */
    /**
     * Sample code: Get workspaces by subscription.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getWorkspacesBySubscription(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspaces().list(Context.NONE);
    }
}
