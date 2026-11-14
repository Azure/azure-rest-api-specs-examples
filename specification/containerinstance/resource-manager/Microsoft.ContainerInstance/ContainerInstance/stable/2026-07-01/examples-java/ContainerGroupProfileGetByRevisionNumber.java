
/**
 * Samples for CGProfile GetByRevisionNumber.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfileGetByRevisionNumber.json
     */
    /**
     * Sample code: ContainerGroupProfileGetByRevisionNumber.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfileGetByRevisionNumber(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfiles().getByRevisionNumberWithResponse("demo", "demo1", "1",
            com.azure.core.util.Context.NONE);
    }
}
