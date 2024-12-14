
/**
 * Samples for KubernetesVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * ListKubernetesVersions.json
     */
    /**
     * Sample code: ListKubernetesVersions.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        listKubernetesVersions(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.kubernetesVersions().list(
            "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation",
            com.azure.core.util.Context.NONE);
    }
}
