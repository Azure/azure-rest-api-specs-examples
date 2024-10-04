
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void operationList(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
