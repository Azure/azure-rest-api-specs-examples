
/**
 * Samples for VmSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ListAvailableContainerServiceVmSkus.json
     */
    /**
     * Sample code: Lists all available Container Service VM SKUs for a location.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listsAllAvailableContainerServiceVMSKUsForALocation(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getVmSkus().list("westus", null, com.azure.core.util.Context.NONE);
    }
}
