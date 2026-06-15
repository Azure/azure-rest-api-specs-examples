
/**
 * Samples for VmSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ListAvailableContainerServiceVmSkusWithExtendedLocations.json
     */
    /**
     * Sample code: Lists all available Container Service VM SKUs with Extended Location information.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listsAllAvailableContainerServiceVMSKUsWithExtendedLocationInformation(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getVmSkus().list("westus", true, com.azure.core.util.Context.NONE);
    }
}
