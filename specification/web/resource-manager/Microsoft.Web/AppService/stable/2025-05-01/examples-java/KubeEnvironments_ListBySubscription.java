
/**
 * Samples for KubeEnvironments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/KubeEnvironments_ListBySubscription.json
     */
    /**
     * Sample code: List kube environments by subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listKubeEnvironmentsBySubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getKubeEnvironments().list(com.azure.core.util.Context.NONE);
    }
}
