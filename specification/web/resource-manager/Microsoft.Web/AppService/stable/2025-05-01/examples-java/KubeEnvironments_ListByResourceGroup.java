
/**
 * Samples for KubeEnvironments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/KubeEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: List kube environments by resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listKubeEnvironmentsByResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getKubeEnvironments().listByResourceGroup("examplerg",
            com.azure.core.util.Context.NONE);
    }
}
