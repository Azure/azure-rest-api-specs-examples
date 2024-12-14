
import com.azure.resourcemanager.hybridcontainerservice.fluent.models.KubernetesVersionProfileInner;
import com.azure.resourcemanager.hybridcontainerservice.models.ExtendedLocation;
import com.azure.resourcemanager.hybridcontainerservice.models.ExtendedLocationTypes;

/**
 * Samples for ResourceProvider PutKubernetesVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * PutKubernetesVersions.json
     */
    /**
     * Sample code: PutKubernetesVersions.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        putKubernetesVersions(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.resourceProviders().putKubernetesVersions(
            "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation",
            new KubernetesVersionProfileInner()
                .withExtendedLocation(new ExtendedLocation().withType(ExtendedLocationTypes.CUSTOM_LOCATION).withName(
                    "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation")),
            com.azure.core.util.Context.NONE);
    }
}
