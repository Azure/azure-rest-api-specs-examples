
/**
 * Samples for Grafana CheckEnterpriseDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/EnterpriseDetails_Post.json
     */
    /**
     * Sample code: EnterpriseDetails_Post.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void enterpriseDetailsPost(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().checkEnterpriseDetailsWithResponse("myResourceGroup", "myWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
