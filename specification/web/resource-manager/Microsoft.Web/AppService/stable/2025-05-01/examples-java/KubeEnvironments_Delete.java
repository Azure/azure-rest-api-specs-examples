
/**
 * Samples for KubeEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/KubeEnvironments_Delete.json
     */
    /**
     * Sample code: Delete kube environment by name.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteKubeEnvironmentByName(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getKubeEnvironments().delete("examplerg", "examplekenv",
            com.azure.core.util.Context.NONE);
    }
}
