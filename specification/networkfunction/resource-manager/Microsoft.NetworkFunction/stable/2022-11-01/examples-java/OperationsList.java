
/**
 * Samples for NetworkFunction ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/
     * OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void operationsList(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.networkFunctions().listOperations(com.azure.core.util.Context.NONE);
    }
}
