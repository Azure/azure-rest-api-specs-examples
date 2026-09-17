
/**
 * Samples for KubernetesVersions GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/KubernetesVersions_Get.json
     */
    /**
     * Sample code: Get Kubernetes versions.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getKubernetesVersions(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesVersions().getByResourceGroupWithResponse("resourceGroupName", "default",
            com.azure.core.util.Context.NONE);
    }
}
