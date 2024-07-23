
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurefleet/AzureFleet.Management/examples/2024-05-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void operationsList(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
