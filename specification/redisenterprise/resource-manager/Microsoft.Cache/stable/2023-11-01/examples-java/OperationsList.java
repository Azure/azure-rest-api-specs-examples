
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void operationsList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
