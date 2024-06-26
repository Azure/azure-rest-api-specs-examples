/** Samples for HybridContainerService ListOrchestrators. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ListOrchestrators.json
     */
    /**
     * Sample code: ListOrchestrators.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listOrchestrators(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .hybridContainerServices()
            .listOrchestratorsWithResponse(
                "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation",
                com.azure.core.util.Context.NONE);
    }
}
