import com.azure.core.util.Context;

/** Samples for NetworkFunction ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void operationsList(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.networkFunctions().listOperations(Context.NONE);
    }
}
