
/**
 * Samples for KubeEnvironments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/KubeEnvironments_Get.json
     */
    /**
     * Sample code: Get kube environments by name.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getKubeEnvironmentsByName(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getKubeEnvironments().getByResourceGroupWithResponse("examplerg", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
