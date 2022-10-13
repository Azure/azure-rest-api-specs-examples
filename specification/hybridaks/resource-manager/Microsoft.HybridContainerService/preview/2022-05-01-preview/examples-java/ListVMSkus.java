import com.azure.core.util.Context;

/** Samples for HybridContainerService ListVMSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListVMSkus.json
     */
    /**
     * Sample code: ListVMSkus.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listVMSkus(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .hybridContainerServices()
            .listVMSkusWithResponse(
                "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation",
                Context.NONE);
    }
}
